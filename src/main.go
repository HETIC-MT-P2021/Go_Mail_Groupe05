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
		err := auth.TokenIsValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"success": false,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func userCredentialsAreValid(email string, password string) bool {
	return email == "akakpo.jeanjacques@gmail.com" && password == "azerty"
}

func getFieldInBody(c *gin.Context, fieldName string) string {
	mapValues := map[string]string{}
	if err := c.ShouldBindJSON(&mapValues); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return ""
	}

	return mapValues[fieldName]
}

func attemptLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if !userCredentialsAreValid(email, password) {
		c.JSON(http.StatusOK, gin.H{
			"tokens":  false,
			"success": false,
			"message": "Please provide valid login credentials",
		})
	} else {
		tokens, _ := auth.GenerateToken(email + password)
		c.JSON(http.StatusCreated, gin.H{
			"tokens": map[string]string{
				"access_token":  tokens.AccessToken,
				"refresh_token": tokens.RefreshToken,
			},
			"message": "Logged in successfully",
			"success": true,
		})
	}
}

func refreshToken(c *gin.Context) {
	refreshToken := getFieldInBody(c, "refresh_token")
	userID, err := auth.RefreshTokenIsValid(refreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"tokens":  false,
			"message": err.Error(),
			"success": false,
		})
	} else {
		tokens, _ := auth.GenerateToken(userID)
		c.JSON(http.StatusCreated, gin.H{
			"tokens": map[string]string{
				"access_token":  tokens.AccessToken,
				"refresh_token": tokens.RefreshToken,
			},
			"message": "Tokens refreshed",
			"success": true,
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
		public.POST("/refresh-token", refreshToken)
	}

	api := router.Group("/api")
	api.Use(authMiddleware())
	{
		api.GET("/users", isRunning)
	}

	router.Run(":" + env["API_PORT"])
}
