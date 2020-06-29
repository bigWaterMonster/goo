package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func main() {
	file, err := os.Open("./1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	f, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println(f.Size())
	r := bufio.NewReader(file)
	s := make([]byte, 7)
	for {
		n, err := r.Read(s)
		fmt.Println(n, err)
		if err != nil {
			break
		}
	}
}
