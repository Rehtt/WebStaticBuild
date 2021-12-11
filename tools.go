/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/12/11 上午 9:27
 */

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
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
