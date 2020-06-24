package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

func dede(c *gin.Context) {
	c.HTML(http.StatusOK, "hh", gin.H{
		"title": "Main website",
	})
}

func jsonp(c *gin.Context) {
	data := gin.H{
		"foo": "bar",
	}
	c.JSONP(http.StatusOK, data)
}

var html = template.Must(template.New("hh").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

type LoginForm struct {
	User     string `form:"user" binding:"required" json:"u"`
	Password string `form:"password" binding:"required" json:"p"`
}

// 表单绑定结构体
func login(c *gin.Context) {
	var form LoginForm
	var form2 LoginForm
	fmt.Println(c. ContentType())
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		fmt.Println(err)
		c.JSON(401, gin.H{"status": err.Error()})
		return
	}
	fmt.Println(form.User, form.Password)
	if c.ShouldBindBodyWith(&form2, binding.JSON) != nil {
		c.JSON(401, gin.H{"status": "form2 unauthorized"})
		return
	}
	fmt.Println(form2.User, form2.Password)
	c.JSON(200, gin.H{"status": "you are logged in"})
}

// 直接获得表单字段值
func formPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	fmt.Println(message, nick)
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func pureJson(c *gin.Context) {
	c.PureJSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}

func main() {
	r := gin.Default()
	// 自定义静态资源路径
	r.Static("/assets", "./assets")
	// 自定义分隔符
	r.Delims("<%", "%>")
	// 表单内存大小
	r.MaxMultipartMemory = 8 << 20
	// 加载模版
	r.LoadHTMLGlob("templates/*")
	// 直接设置模版
	// r.SetHTMLTemplate(html)
	// 路由
	r.GET("/index", index)
	r.GET("/dede", dede)
	r.GET("/JSONP", jsonp)
	r.POST("/login", login)
	r.POST("/form_post", formPost)
	r.GET("/purejson", pureJson)
	r.POST("/post", func(c *gin.Context) {
		c.GetQuery("id")
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.DefaultPostForm("message", "seeyou")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
		})
	})
	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})
	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		// file, _ := c.FormFile("file")
		// 多文件
		form, _ := c.MultipartForm()
		f1 := form.File["file"]
		for index,f  := range f1 {
			fmt.Println(index, f.Filename)
		}
		f2 := form.File["ff"]
		for index,f  := range f2 {
			fmt.Println(index, f.Filename)
		}
		// 上传文件至指定目录
		// c.SaveUploadedFile(file, "./ff")
		// c.SaveUploadedFile(f2, "./ffff")
		// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", f1.Filename))
	})
	r.GET("/someDataFromReader", func(c *gin.Context) {
		res, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || res.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := res.Body
		length := res.ContentLength
		contentType := res.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, length, contentType, reader, extraHeaders)
	})
	// 服务监听
	// r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
	endless.ListenAndServe(":8080", r)
}
