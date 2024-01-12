package main

import (
	"sistema/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.SetupRouter(router)
	router.Run(":8080")
}
