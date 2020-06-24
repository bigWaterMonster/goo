package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles("./index.tmpl")
		if err != nil {
			fmt.Println("ping", err)
			writer.Write([]byte("出错了"))
			return
		}
		hobbyList := []string{"唱", "跳", "rap", "篮球"}
		// user := User{"小鸡吧", "男", 26}
		m1 := map[string]interface{}{
			"name":   "小鸡吧",
			"gender": "男",
			"age":    "22",
			"hobby":  hobbyList,
		}
		err = t.Execute(writer, m1)
		if err != nil {
			fmt.Println("鸡吧出错")
		}
	})
	http.HandleFunc("/pong", func(writer http.ResponseWriter, request *http.Request) {
		// 定义 解析 渲染
		kua := func(name string) (string, error) {
			return name + "吃屎", nil
		}
		t := template.New("second.tmpl")
		t.Funcs(template.FuncMap{
			"kua": kua,
		})
		// 第一个为主，后面为主内嵌套模版
		_, err := t.ParseFiles("./second.tmpl", "./ul.tmpl")
		if err != nil {
			fmt.Println("/pong", err)
			writer.Write([]byte("出错了"))
			return
		}
		name := "超人"

		t.Execute(writer, name)
	})
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles("./base.tmpl", "./home.tmpl")
		if err != nil {
			fmt.Println("ping", err)
			writer.Write([]byte("出错了"))
			return
		}
		msg := "home home"
		t.ExecuteTemplate(writer, "home.tmpl", msg)
	})
	http.HandleFunc("/home2", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles("./base.tmpl", "./home2.tmpl")
		if err != nil {
			fmt.Println("ping", err)
			writer.Write([]byte("出错了"))
			return
		}
		msg := "home23"
		t.ExecuteTemplate(writer, "home2.tmpl", msg)
	})

	http.ListenAndServe(":8080", nil)
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
