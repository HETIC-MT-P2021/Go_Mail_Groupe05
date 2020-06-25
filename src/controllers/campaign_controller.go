package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomail/models"
)

// CreateCampaign handle request to send create a new campaign
func (paramHandler *HandleDb) CreateCampaign(c *gin.Context) {
	dbConnection := paramHandler.DbCon

	campaignName := c.PostForm("campaign_name")
	mailingListID := c.PostForm("mailing_list_id")
	businessID := c.PostForm("business_id")

	thisCampaign, err := models.CreateCampaignWithExistingMailingList(campaignName, mailingListID, businessID, dbConnection)

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
func (paramHandler *HandleDb) CreateCampaignAndMailingList(c *gin.Context) {
	dbConnection := paramHandler.DbCon

	campaignName := c.PostForm("campaign_name")
	mailingListName := c.PostForm("mailing_list_name")
	businessID := c.PostForm("business_id")

	thisMailingList, thisCampaign, err := models.CreateCampaignAndMailingList(mailingListName, campaignName, businessID, dbConnection)

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
func (paramHandler *HandleDb) GetCampaign(c *gin.Context) {
	dbConnection := paramHandler.DbCon

	campaignID := c.PostForm("campaign_id")

	thisCampaign, err := models.GetCampaign(campaignID, dbConnection)

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
func (paramHandler *HandleDb) GetCampaignByBusinessID(c *gin.Context) {
	dbConnection := paramHandler.DbCon

	businessID := c.PostForm("business_id")

	thisCampaigns, err := models.GetBusinessCampaigns(businessID, dbConnection, c)

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
