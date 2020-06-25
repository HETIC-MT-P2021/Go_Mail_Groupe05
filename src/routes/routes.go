package routes

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomail/controllers"
)

func isRunning(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API is running successfully",
		"success": true,
	})
}

// StartRouter will launch the web server
func StartRouter(apiPort string, dbCon *sql.DB, saltString string) {
	router := gin.New()

	public := router.Group("/")
	{
		public.GET("/", isRunning)

		Obj := new(controllers.HandleDbSalt)
		Obj.Db = dbCon
		Obj.SaltString = saltString

		public.POST("/login", Obj.AttemptLogin)

		public.POST("/refresh-token", controllers.RefreshToken)
	}

	router.Run(":" + apiPort)
}
