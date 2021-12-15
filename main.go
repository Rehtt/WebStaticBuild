/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/12/11 上午 9:24
 */

package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

var (
	routerFile   = "router.go"
	resourceFile = "resource.go"
	inFile       = flag.String("path", "./", "输入文件路径")
	outPath      = flag.String("out", "./", "输出路径")
	goPackage    = flag.String("package", "", "go package名，默认为输出路径文件夹名")
	gz           = flag.Bool("gzip", true, "开启gzip预压缩")
	gzLevel      = flag.Int("gz_level", 9, "gzip压缩比")
	printInfo    = flag.Bool("print", false, "输出详情")
	excludeExt   = flagM{}
	excludeFile  = flagM{}
	//excludeFile  = map[string]struct{}{}
	//excludeExt   = flag.String("e_ext", "", "排除录入的文件后缀，多个后缀用|分隔（-e_ext .go|.jpg）")
	//excludeFile  = flag.String("e_file", "", "排除录入的文件，多个用|分隔（-e_file ./file1|./web/file2）")
)

type flagM map[string]struct{}

func (e flagM) String() string {
	return fmt.Sprintf("%v", map[string]struct{}(e))
}
func (e flagM) Set(value string) error {
	e[value] = struct{}{}
	return nil
}

func main() {
	flag.Var(&excludeExt, "e_ext", "排除录入的文件后缀，多个后缀多写几次（ WebStaticBuild -path ./ -e_ext .exe -e_ext .go ）")
	flag.Var(&excludeFile, "e_file", "排除录入的文件，多个文件多写几次（ WebStaticBuild -path ./ -e_file ./file.go -e_file ./web/test.js ）")
	flag.Parse()
	// 转化为绝对地址
	for k := range excludeFile {
		v, err := filepath.Abs(k)
		delete(excludeFile, k)
		if err != nil {
			continue
		}
		excludeFile[v] = struct{}{}
	}

	routerFile = *outPath + "/" + routerFile
	resourceFile = *outPath + "/" + resourceFile

	inFilesPath, err := GetAllFiles(*inFile)
	if err != nil {
		LogErr(err)
	}

	if *goPackage == "" {
		*goPackage, err = GetDirName(*outPath)
		if err != nil {
			LogErr(err)
		}
	}
	router, err := NewGoFile(routerFile, *goPackage)
	if err != nil {
		LogErr(err)
	}
	defer router.Close()

	resource, err := NewGoFile(resourceFile, *goPackage)
	if err != nil {
		LogErr(err)
	}
	defer resource.Close()
	goFileNames := make([]string, len(inFilesPath))
	isGzip := make([]bool, len(inFilesPath))
	for i, v := range inFilesPath {
		goFileNames[i] = "res" + MD5(v)
		isGzip[i] = *gz
		resource.WriteVarBytesToGoFile(*inFile+v, goFileNames[i], *gz, *gzLevel)
		if *printInfo {
			fmt.Println(v)
		}
	}
	router.WriteRouterToGoFile(inFilesPath, goFileNames, isGzip)
}
