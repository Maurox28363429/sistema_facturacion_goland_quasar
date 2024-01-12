package routers

import (
	"fmt"
	"proyecto/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.GET("/users", middlewares.Token, func(c *gin.Context) {
		fmt.Println("Get users")
	})
	router.GET("/users/:id", middlewares.Token, func(c *gin.Context) {
		fmt.Println("Get users")
	})
	router.POST("/users", middlewares.Token, func(c *gin.Context) {
		fmt.Println("Get users")
	})
	router.PUT("/users/:id", middlewares.Token, func(c *gin.Context) {
		fmt.Println("Get users")
	})
	router.DELETE("/users/:id", middlewares.Token, func(c *gin.Context) {
		fmt.Println("Get users")
	})
}
