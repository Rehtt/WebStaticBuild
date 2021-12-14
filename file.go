/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/12/11 上午 9:25
 */

package main

import (
	"fmt"
	"github.com/Rehtt/GoTools"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type write struct {
	f *os.File
}

func NewGoFile(path, goPackage string) (w write, err error) {
	out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return w, err
	}
	out.WriteString("/*此文件为自动生成文件，不要直接做出修改。如果要修改应该修改源文件再执行生成。*/\n\n")
	out.WriteString("package " + goPackage)
	w = write{out}
	return w, err
}
func (w write) Close() error {
	return w.f.Close()
}

func (w *write) WriteString(data string) (int, error) {
	return w.f.WriteString(data)
}
func (w write) WriteVarBytesToGoFile(path, goName string, gz bool, gzLevel int) error {
	//pathFile, err := os.Open(path)
	//if err != nil {
	//	return err
	//}
	w.WriteString(fmt.Sprintf("\nvar %s = []byte{", goName))
	//var writeData bytes.Buffer
	//data := make([]byte, 1024)
	//offset := int64(0)
	//for {
	//	n, err := pathFile.ReadAt(data, offset)
	//	offset += int64(n)
	//	if err != nil && err != io.EOF {
	//		return err
	//	}
	//	if n == 0 {
	//		break
	//	}
	//	for i := range data[:n] {
	//		writeData.WriteString(strconv.FormatUint(uint64(data[i]), 10))
	//		if !gz {
	//			writeData.WriteByte(',')
	//		}
	//	}
	//}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if gz {
		data, err = GZIP(data, gzLevel)
	}
	w.WriteString(GoTools.JoinToString(data, ","))
	//w.WriteString(writeData.String()[:writeData.Len()-1])
	w.WriteString("}")
	return err
}

func (w write) WriteRouterToGoFile(path, value []string, isGzip []bool) {
	/*
		type RouterInfo struct{
			Data		[]byte
			ContentType	string
			Gzip		bool
		}
	*/
	w.WriteString("\ntype RouterInfo struct{\n\tData\t\t[]byte\n\tContentType\tstring\n\tGzip\t\tbool\n}")

	/*
		func GetRouter(url string) (RouterInfo,bool){
			r, ok := router[url]
			return r, ok
		}
	*/
	w.WriteString("\nfunc GetRouter(url string) (RouterInfo, bool) {\n\tr, ok := router[url]\n\treturn r, ok\n}")

	/*
		var router = map[string]RouterInfo{}
	*/
	w.WriteString("\nvar router = map[string]RouterInfo{")
	for i := range path {
		contentType := GetFileContentType(path[i])
		w.WriteString(fmt.Sprintf("\n\t\"%s\": {\n\t\tData:        %s,\n\t\tContentType: \"%s\",\n\t\tGzip:        %t,\n\t},", path[i], value[i], contentType, isGzip[i]))
	}
	w.WriteString("\n}")
}

//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	//PthSep := string(os.PathSeparator)
	PthSep := "/"
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			if _, ok := excludeExt[filepath.Ext(fi.Name())]; !ok {
				path := dirPth + PthSep + fi.Name()
				p, _ := filepath.Abs(path)
				if _, ok = excludeFile[p]; !ok {
					files = append(files, strings.TrimPrefix(path, "./"))
				}
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			t, _ := filepath.Abs(temp1)
			if _, ok := excludeFile[t]; !ok {
				files = append(files, strings.TrimPrefix(temp1, "./"))
			}

		}
	}

	return files, nil
}
