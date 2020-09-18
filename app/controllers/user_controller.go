package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/models"
)

// CreateUser handle request to send create a new user
func CreateUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	businessID := c.PostForm("businessID")

	newUser, err := models.CreateUser(email, password, businessID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Created user successfully",
			"success": true,
			"content": newUser,
		})
	}
}
