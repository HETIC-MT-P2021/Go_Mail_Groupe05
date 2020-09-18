package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomail/models"
)

// CreateBusiness handle request to send create a new user
func CreateBusiness(c *gin.Context) {
	businessName := c.PostForm("businessName")

	newBusiness, err := models.CreateBusiness(businessName)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Created business successfully",
			"success": true,
			"content": newBusiness,
		})
	}
}

// GetBusiness handle request to get new business
func GetBusiness(c *gin.Context) {
	businessID := c.Param("businessID")

	thisBusiness, err := models.GetBusiness(businessID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":     false,
			"message":     err,
			"createdUser": false,
		})
	} else {
		c.JSON(http.StatusFound, gin.H{
			"message":     "Found business successfully",
			"success":     true,
			"createdUser": thisBusiness,
		})
	}
}
