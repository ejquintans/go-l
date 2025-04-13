package user

import (
	"github.com/ejquintans/go-l/internal/domain"
	"github.com/gin-gonic/gin"
	"log"
)

func (h UserHandler) CreateUser(c *gin.Context) {
	var userCreateParams domain.User
	if err := c.BindJSON(&userCreateParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID, err := h.UserService.Create(userCreateParams)
	if err != nil {
		log.Println("Error al insertar usuario:", err)
		c.JSON(500, gin.H{"error": "Error al insertar usuario"})
		return
	}

	c.JSON(200, gin.H{"user_id": userID})
}
