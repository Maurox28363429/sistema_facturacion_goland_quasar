package middleware

import (
	"github.com/gin-gonic/gin"
)

func jwt(c *gin.Context) {
	// Código del middleware
	// ...
	// Llama a c.Next() para continuar con la ejecución del siguiente handler
	c.Next()
}
