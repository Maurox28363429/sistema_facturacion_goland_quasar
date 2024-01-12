package routers

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.GET("/users", token, handlers.getUsers)
	router.GET("/users/:id", token, handlers.getUser)
	router.POST("/users", token, handlers.createUser)
	router.PUT("/users/:id", token, handlers.updateUser)
	router.DELETE("/users/:id", token, handlers.deleteUser)
}
