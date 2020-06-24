package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// 提交表单 a=1&b=2编码后放入body
func PostForm() {
	data := url.Values{}
	data.Add("name", "hello world")
	payload := data.Encode()
	// 字符串reader
	res, _ := http.Post("http://httpbin.org/post", "application/x-www-form-urlencoded", strings.NewReader(payload))
	PrintBody(res)
}

// 提交json 结构体做json
type User struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}

func PostJson() {
	u := User{"陈昊", 18}
	payload, _ := json.Marshal(u)
	// 字节reader
	res, _ := http.Post("http://httpbin.org/post", "application/json", bytes.NewReader(payload))
	PrintBody(res)
}

func main() {
	PostJson()
}
