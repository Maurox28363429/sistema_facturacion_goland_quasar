package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/users", func(c *gin.Context) {
		fmt.Println("Get users")
	})
	router.GET("/users/:id", func(c *gin.Context) {
		fmt.Println("Get users")
	})
	router.POST("/users", func(c *gin.Context) {
		fmt.Println("Get users")
	})
	// ... otras rutas
}
