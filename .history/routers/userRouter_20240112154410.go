package routers

import (
	"proyecto/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.GET("/users", middlewares.Token, func(c *gin.Context) {

	})
	router.GET("/users/:id", middlewares.Token, func(c *gin.Context) {

	})
	router.POST("/users", middlewares.Token, func(c *gin.Context) {

	})
	router.PUT("/users/:id", middlewares.Token, func(c *gin.Context) {

	})
	router.DELETE("/users/:id", middlewares.Token, func(c *gin.Context) {

	})
}
