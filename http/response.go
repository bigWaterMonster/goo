package main

import (
	"bufio"
	"fmt"
	"net/http"

	"golang.org/x/net/html/charset"
)

// 获取状态吗
func status(res *http.Response)  {
	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)
}

// 获取响应头
func header(res *http.Response)  {
	fmt.Println(res.Header.Get("Content-type"))
}

// html 编码
func encoding(res *http.Response)  {
	bufReader := bufio.NewReader(res.Body)
	bytes, _ := bufReader.Peek(1024)
	e, name, ok := charset.DetermineEncoding(bytes, res.Header.Get("Content-type"))
	fmt.Println(e, name, ok)
}
