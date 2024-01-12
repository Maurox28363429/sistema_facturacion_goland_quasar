package routers

import (
	"facturacion/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/users", handlers.GetUserList)
	router.GET("/users/:id", handlers.GetUserByID)
	router.POST("/users", handlers.CreateUser)
	// ... otras rutas
}
