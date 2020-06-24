package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"packages.hetic.net/gomail/auth"
	"packages.hetic.net/gomail/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type handleDbSalt struct {
	Db         *sql.DB
	SaltString string
}

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

func (paramHandler *handleDbSalt) attemptLogin(c *gin.Context) {
	dbConnection := paramHandler.Db
	saltString := paramHandler.SaltString

	email := c.PostForm("email")
	password := c.PostForm("password")

	if !db.VerifyUserCredentials(email, password, dbConnection, saltString) {
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
	refreshToken := c.PostForm("refresh_token")

	fmt.Println(refreshToken)
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
	env, _ := godotenv.Read(".env")

	dbPort, err := strconv.ParseInt(env["DB_PORT"], 10, 64)

	if err != nil {
		panic(err)
	}

	var dbCon = db.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], dbPort)

	router := gin.New()

	public := router.Group("/")
	{
		public.GET("/", isRunning)

		Obj := new(handleDbSalt)
		Obj.Db = dbCon
		Obj.SaltString = env["PW_SALT"]

		public.POST("/login", Obj.attemptLogin)

		public.POST("/refresh-token", refreshToken)
	}

	router.Run(":" + env["API_PORT"])
}
