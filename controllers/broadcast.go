package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/models"

)

// BroadcastCampaign handle request to send a mail to all customer of a campaign
func BroadcastCampaign(c *gin.Context) {
	campaignID := c.PostForm("campaignID")
	_ = c.PostForm("mailFrom")
	_ = c.PostForm("mailContent")
	_ = c.PostForm("mailSubject")

	campaign, err := models.GetCampaign(campaignID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})

		return
	}

	_, customers, err := models.GetMailingList(fmt.Sprintf("%d", campaign.MailingListID), c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
		return
	}

	var customerEmails []string

	for c := 0; c < len(customers); c++ {
		customerEmails = append(customerEmails, customers[c].Email)
	}



	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusFound, gin.H{
			"success": true,
			"message": "Send mail successfully",
			"content": false,
		})
	}
}
