package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "用户未登陆"})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

func GetTopicDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id": c.Param("topic_id"),
	})
}

func NewTopic(c *gin.Context)  {
	c.String(http.StatusOK, "新增")
}

func DelTopic(c *gin.Context)  {
	c.String(http.StatusOK, "删除")
}

func GetTopicList(c *gin.Context) {
	// c.String()
}
