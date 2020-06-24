package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(context *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	context.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}

func Success(context *gin.Context, data gin.H, msg string) {
	Response(context, http.StatusOK, 200, data, msg)
}

func Fail(context *gin.Context, data gin.H, msg string) {
	Response(context, http.StatusOK, 400, data, msg)
}
