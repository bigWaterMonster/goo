package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func rrCookie()  {
	// url := "http://httpbin.org/cookies/set?name=chenhao&age=19"
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(http.MethodGet,"http://httpbin.org/cookies/set?name=chenhao&age=19", nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	r, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	cookies := res.Cookies()
	for _, v := range cookies {
		fmt.Println(v)
	}
	fmt.Println("body", string(r))
}

