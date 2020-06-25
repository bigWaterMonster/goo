package main

import (
	"errors"
	"fmt"
	"net/http"
)

func redirectLimitTimes() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(len(via))
			if len(via) > 10 {
				return errors.New("redirect too times")
			}
			return nil
		},
	}
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/redirect/20", nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	fmt.Println(res.StatusCode)
	if err != nil {
		panic(err)
	}
	PrintBody(res)
}

func redirectForbidden() {

}

// func main() {
// 	redirectLimitTimes()
// }
