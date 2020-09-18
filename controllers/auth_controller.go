package controllers

import (
	"net/http"

	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/models"
	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware Verify the token is valid for secured routes
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenIsValid(c.Request)
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

// AttemptLogin handle request to send an accessToken for a given user
func AttemptLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	isGoodPassword, err := models.VerifyUserCredentials(email, password)

	if !isGoodPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"content": false,
			"success": false,
			"message": err,
		})
	} else {
		tokens, _ := utils.GenerateToken(email + password)
		c.JSON(http.StatusCreated, gin.H{
			"content": map[string]string{
				"access_token":  tokens.AccessToken,
				"refresh_token": tokens.RefreshToken,
			},
			"message": "Logged in successfully",
			"success": true,
		})
	}
}

// RefreshToken handle request to send a new accessToken
func RefreshToken(c *gin.Context) {
	refreshToken := c.PostForm("refresh_token")

	userID, err := utils.RefreshTokenIsValid(refreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"tokens":  false,
			"message": err.Error(),
			"success": false,
		})
	} else {
		tokens, _ := utils.GenerateToken(userID)
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
