package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	router.GET("/api/auth", func(c *gin.Context) {
		fmt.Println("Get users")
	})
	router.POST("/api/auth", func(c *gin.Context) {
		fmt.Println("Get users")
	})
	// ... otras rutas
}
