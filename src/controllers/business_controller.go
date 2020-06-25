package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomail/models"
)

// CreateBusiness handle request to send create a new user
func (paramHandler *HandleDb) CreateBusiness(c *gin.Context) {
	dbConnection := paramHandler.DbCon

	businessName := c.PostForm("business_name")

	newBusiness, err := models.CreateBusiness(businessName, dbConnection)

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
func (paramHandler *HandleDb) GetBusiness(c *gin.Context) {
	dbConnection := paramHandler.DbCon

	businessID := c.PostForm("business_id")

	thisBusiness, err := models.GetBusiness(businessID, dbConnection)

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

// GetBusinessByName handle request to get a business by it's name
func (paramHandler *HandleDb) GetBusinessByName(c *gin.Context) {
	dbConnection := paramHandler.DbCon

	businessName := c.PostForm("business_name")

	thisBusiness, err := models.GetBusinessByName(businessName, dbConnection)

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
