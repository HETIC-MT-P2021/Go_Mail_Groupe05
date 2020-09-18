package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomail/models"
	"packages.hetic.net/gomail/utils"
)

// BroadcastCampaign handle request to send a mail to all customer of a campaign
func BroadcastCampaign(c *gin.Context) {
	campaignID := c.PostForm("campaignID")
	mailFrom := c.PostForm("mailFrom")
	content := c.PostForm("mailContent")
	subject := c.PostForm("mailSubject")

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

	err = utils.SendEmail(customerEmails, []string{}, []string{}, subject, content, "", mailFrom)

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
