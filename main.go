package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	r := gin.Default()
	r.GET("/ping", func1)
	r.POST("/apple", func2)
	r.Run(":8081")

}

// User 结构体定义
type User struct {
	Name  string
	Email string
	Age   int
}

func func1(c *gin.Context) {
	// 1. 定义结构体 确定登陆需要哪些字段 如账号 密码 手机号 邮箱等
	// 2. 新建结构体 u :=User{}
	// 3.  c.ShouldBindJson(&u) 将用户的输入绑定到结构体上
	// 4.  在数据库里面去验证 用户输入的账号是否与数据库里存储的账号密码 一致 如果不对 就返回错误 请重试
	// 5.  验证密码正确之后 jwt 生成token 将用户的id 放到token 里面
	// 6. c.JSON(200,gin.H{"token":  生成的Token})
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.Claims(jwt.RegisteredClaims{
		Issuer:    "",
		Subject:   "",
		Audience:  nil,
		ExpiresAt: nil,
		NotBefore: nil,
		IssuedAt:  nil,
		ID:        "11",
	})).SignedString([]byte("secret"))
	c.JSON(200, gin.H{"11": token})
}

func func2(c *gin.Context) {
	a := User{
		Name:  "",
		Email: "",
	}
	c.ShouldBindJSON(&a)
	//name := c.PostForm("name")
	//email := c.PostForm("email")
	//a.Name = name
	//a.Email = email
	c.JSON(200, gin.H{"data": a})
}
