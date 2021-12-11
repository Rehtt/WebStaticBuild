/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/12/11 上午 9:24
 */

package main

import (
	"flag"
)

var (
	routerFile   = "router.go"
	resourceFile = "resource.go"
	inFile       = flag.String("path", "./web", "输入文件路径")
	outPath      = flag.String("out", "./", "输出路径")
	goPackage    = flag.String("package", "", "go package名，默认为输出路径文件夹名")
	gz           = flag.Bool("gzip", true, "开启gzip预压缩")
	gzLevel      = flag.Int("gz_level", 9, "gzip压缩比")
)

func main() {
	flag.Parse()
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
		resource.WriteVarBytesToGoFile(v, goFileNames[i], *gz, *gzLevel)
	}
	router.WriteRouterToGoFile(inFilesPath, goFileNames, isGzip)
}
