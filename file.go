/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/12/11 上午 9:25
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type write struct {
	f *os.File
}

func NewGoFile(path, goPackage string) (w write, err error) {
	out, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
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
func (w write) WriteVarBytesToGoFile(path, goName string) error {
	pathFile, err := os.Open(path)
	if err != nil {
		return err
	}
	w.WriteString(fmt.Sprintf("\nvar %s = []byte{", goName))
	var writeData strings.Builder
	data := make([]byte, 512)
	offset := int64(0)
	for {
		n, err := pathFile.ReadAt(data, offset)
		offset += int64(n)
		if err != nil && err.Error() != "EOF" {
			return err
		}
		if n == 0 {
			break
		}
		for i := range data[:n] {
			writeData.WriteString(strconv.FormatUint(uint64(data[i]), 10))
			writeData.WriteByte(',')
		}
	}
	w.WriteString(writeData.String()[:writeData.Len()-1])
	w.WriteString("}")
	return err
}

func (w write) WriteVarMapStringBytesToGoFile(key, value []string) {
	w.WriteString("\nvar Router = map[string][]byte{")
	for i := range key {
		w.WriteString(fmt.Sprintf("\n\t\"%s\": %d,", key[i], value[i]))
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
			//ok := strings.HasSuffix(fi.Name(), ".go")
			ok := true
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}
