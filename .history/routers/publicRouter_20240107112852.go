package routers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func PublicRouter(router *gin.Engine) {
	// Ruta estática para servir archivos desde la carpeta "./public"
	router.Static("/public", "./public")

	router.POST("/api/upload", func(c *gin.Context) {
		// Obtener el archivo del formulario
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "No se pudo obtener el archivo",
			})
			return
		}

		// Generar una ruta única para almacenar el archivo
		filename := filepath.Base(file.Filename)
		dst := "./public/" + filename

		// Guardar el archivo en la ruta especificada
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "No se pudo guardar el archivo",
			})
			return
		}

		// Devolver la ruta del archivo
		c.JSON(http.StatusOK, gin.H{
			"ruta": "/public/" + filename,
		})
	})
}
