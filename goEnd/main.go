package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	srv := http.Server{
		Addr: ":8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen %s\n", err)
		}
	}()
	// 中断通道
	quit := make(chan os.Signal)
	// 发送中断信号给中断通道
	signal.Notify(quit, os.Interrupt)
	fmt.Println("阻塞quit")
	<- quit
	fmt.Println("阻塞结束")
	log.Println("Shutdown Server ...")
	// main协程创建context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	// Shutdown 中断main协程
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("server exiting")
}
