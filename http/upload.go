package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func PostFile() {
	// 字节包 类似字符串包
	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)
	writer.WriteField("words", "123")
	// 表单字段名 文件名 多文件要分开写入上传
	upw, _ := writer.CreateFormFile("uploadFile1", "uploadFile1")
	file, _ := os.Open("./http/index.tmpl")
	defer func() { file.Close() }()
	io.Copy(upw, file)
	upw2, _ := writer.CreateFormFile("uploadFile2", "uploadFile2")
	file2, _ := os.Open("./http/http.go")
	defer func() { file2.Close() }()
	io.Copy(upw2, file2)
	writer.Close()
	res, _ := http.Post("http://httpbin.org/post", writer.FormDataContentType(), &body)
	PrintBody(res)
}

// func main() {

	// http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
	// 	tem, err := template.ParseFiles("./http/index.tmpl")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	tem.Execute(writer, map[string]string{"title": "8080"})
	// })
	// http.HandleFunc("/JSONP", func(writer http.ResponseWriter, request *http.Request) {
	// 	cd := request.URL.Query().Get("callback")
	// 	writer.Write([]byte(cd+"(123)"))
	// })
	// http.ListenAndServe(":8080", nil)
// 	PostFile()
// }
