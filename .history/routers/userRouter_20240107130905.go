package routers

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.GET("/users", handlers.getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
}
