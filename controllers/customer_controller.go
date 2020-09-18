package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomail/models"
)

// GetCustomer handle request to get a customer
func GetCustomer(c *gin.Context) {
	cutomerID := c.Param("customerID")

	thisCustomer, err := models.GetCustomer(cutomerID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusFound, gin.H{
			"success": true,
			"message": "Found customer successfully",
			"content": thisCustomer,
		})
	}
}

// CreateCustomer handle request to create a customer
func CreateCustomer(c *gin.Context) {
	email := c.PostForm("email")
	name := c.PostForm("name")
	surname := c.PostForm("surname")
	businessID := c.PostForm("business_id")

	thisCustomer, err := models.CreateCustomer(email, name, surname, businessID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Created customer successfully",
			"content": thisCustomer,
		})
	}
}

// CreateAndLinkCustomer handle request to create a customer and link it to a mailing list
func CreateAndLinkCustomer(c *gin.Context) {
	email := c.PostForm("email")
	name := c.PostForm("name")
	surname := c.PostForm("surname")
	businessID := c.PostForm("businessID")
	mailingListID := c.PostForm("mailingListID")

	thisCustomer, err := models.CreateAndLinkCustomer(email, name, surname, businessID, mailingListID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Created customer and link successfully",
			"content": thisCustomer,
		})
	}
}

// UnlinkCustomerMailingList handle request to unlink a customer and a mailing list
func UnlinkCustomerMailingList(c *gin.Context) {
	customerID := c.PostForm("customer_id")
	mailingListID := c.PostForm("mailing_list_id")

	err := models.UnlinkCustomerMailingList(customerID, mailingListID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"success": true,
			"message": "Unlinked customer and mailing list successfully",
			"content": nil,
		})
	}
}
