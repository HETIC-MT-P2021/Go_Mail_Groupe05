package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/app/models"
)

// CreateMailingList handle request to send create a new mailingList
func CreateMailingList(c *gin.Context) {
	mailingListName := c.PostForm("mailing_list_name")
	businessID := c.PostForm("business_id")

	newMailingList, err := models.CreateMailingList(mailingListName, businessID)

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
func GetMailingList(c *gin.Context) {
	mailingListID := c.Param("mailingListID")

	mailingList, customers, err := models.GetMailingList(mailingListID, c)

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
