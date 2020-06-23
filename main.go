package main

import (
	"net/http"

	"packages.hetic.net/gomail/auth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func isRunning(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "API is running successfully",
	})
}

func isAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func attemptAuth(ctx *gin.Context) {
	username := ctx.Param("username")

	if username != "" {
		validToken, _ := auth.GenerateJWT(username)

		ctx.JSON(http.StatusOK, gin.H{
			"token":   validToken,
			"success": true,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
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
