package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	// Lógica para obtener la lista de usuarios
	users := []string{"Usuario 1", "Usuario 2", "Usuario 3"}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserByID(c *gin.Context) {
	// Obtener el ID del parámetro de la ruta
	userID := c.Param("id")

	// Lógica para obtener un usuario por su ID
	user := "Usuario " + userID

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func CreateUser(c *gin.Context) {
	// Lógica para crear un nuevo usuario

	// Obtener los datos del cuerpo de la solicitud
	type CreateUserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos de usuario inválidos",
		})
		return
	}

	// Crear el usuario en la base de datos u otra lógica

	// Responder con el usuario creado
	user := req.Name

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
