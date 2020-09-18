package controllers

import (
	"fmt"
	"net/http"

	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/app/models"
	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/app/producer"
	"github.com/gin-gonic/gin"
)

// BroadcastCampaign handle request to send a mail to all customer of a campaign
func BroadcastCampaign(c *gin.Context) {
	campaignID := c.PostForm("campaignID")
	mailFrom := c.PostForm("mailFrom")
	mailContent := c.PostForm("mailContent")
	mailSubject := c.PostForm("mailSubject")

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

	producer.PublishMailData(mailSubject, mailContent, mailFrom, customerEmails)

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
