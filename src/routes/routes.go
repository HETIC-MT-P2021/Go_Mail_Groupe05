package routes

import (
	"database/sql"
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
func StartRouter(apiPort string, dbCon *sql.DB, saltString string) {
	router := gin.New()

	DbAndSaltHandler := new(controllers.HandleDbAndSalt)
	DbAndSaltHandler.DbCon = dbCon
	DbAndSaltHandler.SaltString = saltString

	DbHandler := new(controllers.HandleDb)
	DbHandler.DbCon = dbCon

	public := router.Group("/")
	{
		public.GET("/", healthCheck)

		public.POST("/login", DbAndSaltHandler.AttemptLogin)
		public.POST("/refresh-token", controllers.RefreshToken)
	}

	apiRoutes := router.Group("/api")
	apiRoutes.Use(controllers.AuthMiddleware())
	{
		// Business

		apiRoutes.GET("/business/:businessID", DbHandler.GetBusiness)

		apiRoutes.POST("business/", DbHandler.CreateBusiness)

		// User

		apiRoutes.POST("/user", DbAndSaltHandler.CreateUser)

		// Campaign

		apiRoutes.GET("/campaign/withid/:campaignID", DbHandler.GetCampaign)
		apiRoutes.GET("/campaign/withbusiness/:businessID", DbHandler.GetCampaignByBusinessID)

		apiRoutes.POST("/campaign", DbHandler.CreateCampaign)
		apiRoutes.POST("/campaign/mailing-list", DbHandler.CreateCampaignAndMailingList)

		// Mailing list

		apiRoutes.GET("/mailing-list/:mailingListID", DbHandler.GetMailingList)

		apiRoutes.POST("/mailing-list", DbHandler.CreateMailingList)

		// Customer

		apiRoutes.GET("/customer/:customerID", DbHandler.GetCustomer)

		apiRoutes.POST("/customer", DbHandler.CreateCustomer)
		apiRoutes.POST("/customer/link/", DbHandler.CreateAndLinkCustomer)
		apiRoutes.POST("/customer/unlink/", DbHandler.UnlinkCustomerMailingList)
	}

	router.Run(":" + apiPort)
}
