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

func authMiddleware() gin.HandlerFunc {
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

func verifyUserCredentials(email string, password string) bool {
	return email == "akakpo.jeanjacques@gmail.com" && password == "azerty"
}

func attemptLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if !verifyUserCredentials(email, password) {
		c.JSON(http.StatusOK, gin.H{
			"token":   false,
			"success": false,
			"message": "Please provide valid login credentials",
		})
	} else {
		token, _ := auth.GenerateToken(email + password)
		c.JSON(http.StatusOK, gin.H{
			"access_token": token,
			"message":      "Logged in successfully",
			"success":      true,
		})
	}
}

func main() {
	env, _ := godotenv.Read()
	router := gin.New()

	public := router.Group("/")
	{
		public.GET("/", isRunning)
		public.POST("/login", attemptLogin)
	}

	api := router.Group("/api")
	api.Use(authMiddleware())
	{
		api.GET("/users", isRunning)
	}

	router.Run(env["API_PORT"])
}
