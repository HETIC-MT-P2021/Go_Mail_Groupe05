package authcontroller

import (
	"database/sql"
	"net/http"

	usermodel "packages.hetic.net/gomail/models/user"
	"packages.hetic.net/gomail/utils/auth"

	"github.com/gin-gonic/gin"
)

// HandleDbSalt is a structure to pass parameters in routes/controllers
type HandleDbSalt struct {
	Db         *sql.DB
	SaltString string
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

// AttemptLogin handle request to send an accessToken for a given user
func (paramHandler *HandleDbSalt) AttemptLogin(c *gin.Context) {
	dbConnection := paramHandler.Db
	saltString := paramHandler.SaltString

	email := c.PostForm("email")
	password := c.PostForm("password")

	if !usermodel.VerifyUserCredentials(email, password, dbConnection, saltString) {
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

// RefreshToken handle request to send a new accessToken
func RefreshToken(c *gin.Context) {
	refreshToken := c.PostForm("refresh_token")

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
