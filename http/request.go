package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PrintBody(res *http.Response)  {
	defer func() {
		res.Body.Close()
	}()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(content))
}

// 加查询参数
func requestByQuery()  {
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	// get请求参数
	queryValues := url.Values{}
	queryValues.Add("name", "chenhao")
	queryValues.Add("age", "18")
	req.URL.RawQuery = queryValues.Encode()
	//
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	PrintBody(res)
}

// 加头

func requestByHead()  {
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	// 加头
	req.Header.Add("user-agent", "chrome")
	//
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	PrintBody(res)
}
