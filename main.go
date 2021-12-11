/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/12/11 上午 9:24
 */

package main

import "flag"

var (
	routerFile   = "router.go"
	resourceFile = "resource.go"
	inFile       = flag.String("path", "./web", "输入文件路径")
	outPath      = flag.String("out", "./", "输出路径")
	goPackage    = flag.String("package", "", "go package名，默认为输出路径文件夹名")
)

func main() {
	flag.Parse()
	routerFile = *outPath + routerFile
	resourceFile = *outPath + resourceFile

	inFilesPath, err := GetAllFiles(*inFile)
	if err != nil {
		LogErr(err)
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
	for i, v := range inFilesPath {
		goFileNames[i] = MD5(v)
		resource.WriteVarBytesToGoFile(v, goFileNames[i])
	}
	router.WriteVarMapStringBytesToGoFile(inFilesPath, goFileNames)
}
