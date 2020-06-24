package main

import (
	"net/http"
)

func get() {
	res, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	PrintBody(res)
}

func post() {
	res, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		panic(err)
	}
	PrintBody(res)
}

func put() {
	req, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	PrintBody(res)
}

func del() {
	req, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	PrintBody(res)
}

