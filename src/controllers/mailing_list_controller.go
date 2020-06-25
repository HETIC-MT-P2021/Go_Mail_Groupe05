package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomail/models"
)

// CreateMailingList handle request to send create a new mailingList
func (paramHandler *HandleDb) CreateMailingList(c *gin.Context) {
	dbConnection := paramHandler.DbCon

	mailingListName := c.PostForm("mailing_list_name")
	businessID := c.PostForm("business_id")

	newMailingList, err := models.CreateMailingList(mailingListName, businessID, dbConnection)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Created Mailing list successfully",
			"success": true,
			"content": newMailingList,
		})
	}
}

// GetMailingList handle request to get all customers of a mailingList
func (paramHandler *HandleDb) GetMailingList(c *gin.Context) {
	dbConnection := paramHandler.DbCon

	mailingListID := c.PostForm("mailing_list_id")

	mailingList, customers, err := models.GetMailingList(mailingListID, dbConnection, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusFound, gin.H{
			"success": true,
			"message": "Found campaign successfully",
			"content": []interface{}{mailingList, customers},
		})
	}
}
