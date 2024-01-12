package routers

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.GET("/users", middleware.token, handlers.getUsers)
	router.GET("/users/:id", middleware.token, handlers.getUser)
	router.POST("/users", middleware.token, handlers.createUser)
	router.PUT("/users/:id", middleware.token, handlers.updateUser)
	router.DELETE("/users/:id", middleware.token, handlers.deleteUser)
}
