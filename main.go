package main

import (
	"net/http"

	"packages.hetic.net/gomail/auth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func isRunning(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API is running successfully",
		"success": true,
	})
}

func isAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		authenticated := auth.VerifyToken(c.Request.Header.Get("Token"))

		if !authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Forbidden! You are not authorized",
				"success": false,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func attemptAuth(c *gin.Context) {
	username := c.Param("username")

	if username != "" {
		validToken, _ := auth.GenerateToken(username)

		c.JSON(http.StatusOK, gin.H{
			"token":   validToken,
			"success": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token":   "",
			"success": false,
			"message": "Please provide a value to :username parameter",
		})
	}
}

func main() {
	env, _ := godotenv.Read()
	router := gin.New()

	public := router.Group("/")
	{
		public.GET("/", isRunning)
		public.POST("/auth/:username", attemptAuth)
	}

	api := router.Group("/api")
	api.Use(isAuthenticated())
	{
		api.GET("/users", isRunning)
	}

	router.Run(env["API_PORT"])
}
