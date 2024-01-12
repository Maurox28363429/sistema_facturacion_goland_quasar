package main

import (
	"os"
	"proyecto/models"
	"proyecto/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//init database
	models.ConnectDB()
	//router
	router := gin.Default()
	routers.UserRouter(router)
	routers.AuthRouter(router)
	routers.PublicRouter(router)
	//config mode
	gin.SetMode(gin.ReleaseMode)
	os.Setenv("GIN_MODE", "release")
	//init gin
	router.Run(":8080")
}
