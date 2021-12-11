/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/12/10 21:19
 */

package main

import (
	"fmt"
	"github.com/Rehtt/WebStaticBuild/example/web"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		r, ok := web.GetRouter("./web" + request.RequestURI)
		if !ok {
			http.Redirect(writer, request, "/index.html#/404", 301)
		}
		writer.Header().Set("Content-Type", r.ContentType)
		if r.Gzip {
			writer.Header().Set("content-encoding", "gzip")
		}
		writer.Write(r.Data)
	})
	fmt.Println("start")
	http.ListenAndServe(":8090", nil)
}
