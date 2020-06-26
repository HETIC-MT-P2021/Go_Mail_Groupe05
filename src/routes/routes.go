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

	business := router.Group("/business")
	{
		business.GET("/:businessID", DbHandler.GetBusiness)
		business.POST("/", DbHandler.CreateBusiness)
	}

	user := router.Group("/user")
	{
		user.POST("/", DbAndSaltHandler.CreateUser)
	}

	campaign := router.Group("/campaign")
	{
		campaign.GET("/:campaignID", DbHandler.GetCampaign)
		campaign.GET("/:businessID", DbHandler.GetCampaignByBusinessID)
		campaign.POST("/", DbHandler.CreateCampaign)
		campaign.POST("/mailing-list", DbHandler.CreateCampaignAndMailingList)

	}

	mailingList := router.Group("/mailing-list")
	{
		mailingList.GET("/:mailingListID", DbHandler.GetMailingList)
		mailingList.POST("/", DbHandler.CreateMailingList)
	}

	customer := router.Group("/customer")
	{
		customer.GET("/:customerID", DbHandler.GetCustomer)
		customer.POST("/", DbHandler.CreateCustomer)
		customer.POST("/link/", DbHandler.CreateAndLinkCustomer)
		customer.POST("/unlink/", DbHandler.UnlinkCustomerMailingList)
	}

	router.Run(":" + apiPort)
}
