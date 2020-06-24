package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Current += int64(n)
	fmt.Printf("\r进度 %.2f", float64(r.Current)/float64(r.Total)*100)
	return n, err
}

// 下载文件带进度条
func Progress(url, filename string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		res.Body.Close()
	}()
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := &Reader{
		Reader: res.Body,
		Total:  res.ContentLength,
	}
	_, err = io.Copy(file, reader)
	if err != nil {
		panic(err)
	}
}

// 下载文件
func downloadFile(url, filename string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		res.Body.Close()
	}()
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	n, err := io.Copy(file, res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

// func main() {
// 	Progress("https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png", "helloWorld.png")
// }
