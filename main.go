package main

import (
	"strconv"

	"github.com/joho/godotenv"
	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/models"
	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/producer"
	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/router"
	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/utils"
)

func main() {
	env, _ := godotenv.Read(".env")

	dbPort, err := strconv.ParseInt(env["DB_PORT"], 10, 64)

	if err != nil {
		panic(err)
	}

	models.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], dbPort)

	utils.InitSMTPCon(env["SMTP_USER"], env["SMTP_PASSWORD"], env["SMTP_HOST"], 25)

	producer.ConnectToRabbit(
		env["RABBIT_HOST"],
		env["RABBIT_PORT"],
		env["RABBIT_USER"],
		env["RABBIT_PASSWORD"],
	)

	router.Configure().Run(":" + env["API_PORT"])
}
