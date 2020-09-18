package main

import (
	"strconv"

	"github.com/joho/godotenv"
	"packages.hetic.net/gomail/models"
	"packages.hetic.net/gomail/routes"
	"packages.hetic.net/gomail/utils"
)

func main() {
	env, _ := godotenv.Read(".env")

	dbPort, err := strconv.ParseInt(env["DB_PORT"], 10, 64)

	if err != nil {
		panic(err)
	}

	models.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], dbPort)

	utils.InitSMTPCon(env["SMTP_USER"], env["SMTP_PASSWORD"])

	routes.StartRouter(env["API_PORT"])
}
