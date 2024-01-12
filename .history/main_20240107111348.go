package main

import (
	"proyecto/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.UserRouter(router)
	router.AuthRouter(router)
	router.Run(":8080")
}
