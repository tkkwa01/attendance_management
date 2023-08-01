package main

import (
	"attendance-management/adopter/gateway/mysql"
	"attendance-management/driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	a := driver.GetRDB
	b := mysql.Chat()
	fmt.Println(b)
	fmt.Println(a)

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":8080")
}
