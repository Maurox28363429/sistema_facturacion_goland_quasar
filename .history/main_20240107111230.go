package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.userRouter(router)
	router.Run(":8080")
}
