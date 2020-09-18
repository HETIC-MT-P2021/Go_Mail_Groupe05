package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/app/models"
)

// CreateCampaign handle request to send create a new campaign
func CreateCampaign(c *gin.Context) {
	campaignName := c.PostForm("campaign_name")
	mailingListID := c.PostForm("mailing_list_id")
	businessID := c.PostForm("business_id")

	thisCampaign, err := models.CreateCampaignWithExistingMailingList(campaignName, mailingListID, businessID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Created campaign successfully",
			"content": thisCampaign,
		})
	}
}

// CreateCampaignAndMailingList handle request to send create a new campaign and a mailing list
func CreateCampaignAndMailingList(c *gin.Context) {
	campaignName := c.PostForm("campaignName")
	mailingListName := c.PostForm("mailingListName")
	businessID := c.PostForm("businessID")

	thisMailingList, thisCampaign, err := models.CreateCampaignAndMailingList(mailingListName, campaignName, businessID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Created campaign successfully",
			"content": []interface{}{thisCampaign, thisMailingList},
		})
	}
}

// GetCampaign handle request to get a campaign
func GetCampaign(c *gin.Context) {
	campaignID := c.Param("campaignID")

	thisCampaign, err := models.GetCampaign(campaignID)

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
			"content": thisCampaign,
		})
	}
}

// GetCampaignByBusinessID handle request to get all campaigns of a business
func GetCampaignByBusinessID(c *gin.Context) {
	businessID := c.Param("businessID")

	thisCampaigns, err := models.GetBusinessCampaigns(businessID, c)

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
			"content": thisCampaigns,
		})
	}
}
