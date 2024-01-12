package main

import (
	"os"
	"proyecto/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.UserRouter(router)
	routers.AuthRouter(router)
	//config mode
	gin.SetMode(gin.ReleaseMode)
	os.Setenv("GIN_MODE", "release")
	//init gin
	router.Run(":8080")
}
