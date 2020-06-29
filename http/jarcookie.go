package main

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func main() {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	res, err := client.Get("http://httpbin.org/cookies/set?username=chenhao&age=321")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

}
