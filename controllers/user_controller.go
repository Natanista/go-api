package controllers

import (
	"github.com/gin-gonic/gin"
	database "github.com/natanista/go-api/database/seeder"
	"github.com/natanista/go-api/models"
	"github.com/natanista/go-api/services"

)

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}
c.Status(204)
	
}