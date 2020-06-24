package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.GET("/topic/:topic_id", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "帖子id为%s", ctx.Param("topic_id"))
	})
	r.Run(":8080")
}
