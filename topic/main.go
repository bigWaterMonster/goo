package main

import (
	"github.com/gin-gonic/gin"

	"g1/topic/src"
)

func main()  {
	r := gin.Default()
	v1 := r.Group("/v1/topics")
	v1.Use(src.MustLogin())
	{
		v1.GET("/:topic_id", src.GetTopicDetail)
		v1.POST("", src.NewTopic)
		v1.DELETE("/:topic_id", src.DelTopic)
	}
	r.Run(":8080")
}
