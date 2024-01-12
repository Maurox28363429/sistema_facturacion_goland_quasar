package routers

import (
	"github.com/gin-gonic/gin"
	"proyecto/middleware"
)

func UserRouter(router *gin.Engine) {
	router.GET("/users",tok, handlers.getUsers)
	router.GET("/users/:id", handlers.getUser)
	router.POST("/users", handlers.createUser)
	router.PUT("/users/:id", handlers.updateUser)
	router.DELETE("/users/:id", handlers.deleteUser)
}
