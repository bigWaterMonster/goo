package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TODO struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var DB *gorm.DB

func initMysql() (err error) {
	DB, err = gorm.Open("mysql", "root:940213@tcp(49.234.156.231:3306)/kkb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	err = DB.DB().Ping()
	if err != nil {
		return err
	}
	DB.AutoMigrate(&TODO{})
	return nil
}

func main() {
	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	r := gin.Default()
	r.LoadHTMLFiles("v/bubble_backend/dist/index.html")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	r.Static("/static", "v/bubble_backend/dist/static")
	todo := r.Group("/v1/todo")
	{
		todo.GET("/", GETALL)
		todo.GET("/:id", GETONE)
		todo.POST("/", POST)
		todo.PUT("/:id", PUT)
		todo.DELETE("/:id", DELETE)
	}
	r.Run(":8080")
}

func GETALL(context *gin.Context) {
	var todoList []TODO
	if err := DB.Debug().Find(&todoList).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	} else {
		context.JSON(http.StatusOK, &todoList)
	}
}

func GETONE(context *gin.Context) {

}

func POST(context *gin.Context) {
	todo := TODO{}
	context.BindJSON(&todo)
	if err := DB.Debug().Create(&todo).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	} else {
		context.JSON(http.StatusOK, &todo)
	}
}

func PUT(context *gin.Context) {
	idStr := context.Param("id")
	id, _ := strconv.Atoi(idStr)
	todo := TODO{ID: id}
	context.BindJSON(&todo)
	if err := DB.Debug().Model(&todo).Update("status", todo.Status).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	} else {
		context.JSON(http.StatusOK, &todo)
	}
}

func DELETE(context *gin.Context) {
	idStr := context.Param("id")
	id, _ := strconv.Atoi(idStr)
	todo := TODO{ID: id}
	if err := DB.Debug().Delete(&todo).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	} else {
		context.JSON(http.StatusOK, &todo)
	}
}
