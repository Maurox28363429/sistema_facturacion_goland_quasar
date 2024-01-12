package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.GET("/api/users", func(c *gin.Context) {
		fmt.Println("Get users")
	})
	router.GET("/api/users/:id", func(c *gin.Context) {
		fmt.Println("Get users")
	})
	router.POST("/api/users", func(c *gin.Context) {
		fmt.Println("Get users")
	})
	// ... otras rutas
}
