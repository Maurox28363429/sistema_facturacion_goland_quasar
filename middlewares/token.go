package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	// Código del middleware
	// ...
	// Llama a c.Next() para continuar con la ejecución del siguiente handler
	c.Next()
}
