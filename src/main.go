package main

import (
	"database/sql"
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

func (paramHandler *handleDbSalt) attemptLogin(c *gin.Context) {
	dbConnection := paramHandler.Db
	saltString := paramHandler.SaltString

	email := c.PostForm("email")
	password := c.PostForm("password")

	if !db.VerifyUserCredentials(email, password, dbConnection, saltString) {
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
	env, _ := godotenv.Read("../.env")

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
	}

	router.Run("0.0.0.0:" + env["API_PORT"])
}
