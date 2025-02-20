package main

import (
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	r := gin.Default()
	r.GET("/ping", func1)
	r.POST("/apple", func2)
	r.Run(":8081")

}
func func1(c *gin.Context) {
	c.JSON(200, "ok")
}
func func2(c *gin.Context) {
	c.JSON(200, "okk")
}
