package main

import (
	"github.com/gin-gonic/gin"
	"prueba.com/m/controllers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/personas", controllers.GetAllPersonas)
	router.POST("/personas", controllers.NewPersona)
	router.Run()
	gin.SetMode(gin.ReleaseMode)
}
