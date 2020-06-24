package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"g1/auth/common"
	"g1/auth/dto"
	"g1/auth/model"
	"g1/auth/response"
	"g1/auth/utils"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	if len(telephone) != 11 {
		// 422 用户格式问题
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		// 422 用户格式问题
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	if utils.IsTelephoneExist(db, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号已存在")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, err.Error())
	}
	newUser := model.User{
		Name:      name,
		Password:  string(hashedPassword),
		Telephone: telephone,
	}
	db.Debug().Create(&newUser)
	response.Response(c, http.StatusOK, 200, gin.H{"id": newUser.ID}, "ok")
}

func Login(c *gin.Context) {
	db := common.GetDB()
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	if len(telephone) != 11 {
		// 422 用户格式问题
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		// 422 用户格式问题
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	user := model.User{}
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, err.Error())
		return
	}
	response.Response(c, http.StatusOK, 200, gin.H{"token": token}, err.Error())
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Response(c, http.StatusOK, 200, gin.H{"user": dto.ToUserDto(user.(*model.User))}, "获取成功")
}
