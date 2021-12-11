/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/12/11 上午 9:27
 */

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

func MD5(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}

func LogErr(err ...interface{}) {
	fmt.Fprint(os.Stderr, err...)
	os.Exit(1)
}

func GZIP(data []byte, level int) ([]byte, error) {
	var b bytes.Buffer
	g, err := gzip.NewWriterLevel(&b, level)
	_, err = g.Write(data)
	if err != nil {
		return nil, err
	}
	g.Close()
	return b.Bytes(), nil
}

func GetFileContentType(path string) string {
	//http.DetectContentType 不准确
	//f, err := os.Open(path)
	//if err != nil {
	//	return "", err
	//}
	//defer f.Close()
	//data := make([]byte, 512)
	//f.Read(data)
	//return http.DetectContentType(data), nil

	return mime.TypeByExtension("." + filepath.Ext(path))
}

func GetDirName(path string) (string, error) {
	p, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	d, _ := filepath.Split(p)
	fmt.Println(d)
	s := strings.Split(d, string(os.PathSeparator))
	if len(s) > 1 {
		return s[len(s)-2], nil
	}
	return s[len(s)-1], nil
}
