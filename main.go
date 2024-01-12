package main

import (
	"log"
	"os"
	"proyecto/models"
	"proyecto/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//inicializar .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}
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
