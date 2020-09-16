package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomail/controllers"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API is running successfully",
		"success": true,
	})
}

// StartRouter will launch the web server
func StartRouter(apiPort string) {
	router := gin.New()

	public := router.Group("/")
	{
		public.GET("/", healthCheck)

		public.POST("/login", controllers.AttemptLogin)
		public.POST("/refresh-token", controllers.RefreshToken)

		// User
		public.POST("/user", controllers.CreateUser)
	}

	apiRoutes := router.Group("/api")
	apiRoutes.Use(controllers.AuthMiddleware())
	{
		// Business
		apiRoutes.GET("/business/:businessID", controllers.GetBusiness)
		apiRoutes.POST("/business", controllers.CreateBusiness)

		// Campaign
		apiRoutes.GET("/campaign/withid/:campaignID", controllers.GetCampaign)
		apiRoutes.GET("/campaign/withbusiness/:businessID", controllers.GetCampaignByBusinessID)
		apiRoutes.POST("/campaign", controllers.CreateCampaign)
		apiRoutes.POST("/campaign/mailing-list", controllers.CreateCampaignAndMailingList)

		// Mailing list
		apiRoutes.GET("/mailing-list/:mailingListID", controllers.GetMailingList)
		apiRoutes.POST("/mailing-list", controllers.CreateMailingList)
		apiRoutes.POST("/broadcast", controllers.BroadcastCampaign)

		// Customer
		apiRoutes.GET("/customer/:customerID", controllers.GetCustomer)
		apiRoutes.POST("/customer", controllers.CreateCustomer)
		apiRoutes.POST("/customer/link/", controllers.CreateAndLinkCustomer)
		apiRoutes.POST("/customer/unlink/", controllers.UnlinkCustomerMailingList)
	}

	router.Run(":" + apiPort)
}
