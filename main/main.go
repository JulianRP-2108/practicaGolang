package main

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", controllers.Ping)
	router.Run()
}
