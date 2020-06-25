package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomail/models"
)

// CreateUser handle request to send create a new user
func (paramHandler *HandleDbAndSalt) CreateUser(c *gin.Context) {
	dbConnection := paramHandler.DbCon
	saltString := paramHandler.SaltString

	email := c.PostForm("email")
	password := c.PostForm("password")
	businessID := c.PostForm("businessID")

	newUser, err := models.CreateUser(email, password, businessID, dbConnection, saltString)

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
