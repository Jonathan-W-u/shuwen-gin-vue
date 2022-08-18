package main

import (
	"shuwen-gin-vue/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	// fmt.Println("shuwen-gin-vue")
	r.Run()
}
