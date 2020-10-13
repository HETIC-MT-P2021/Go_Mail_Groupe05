package main

import (
	"strconv"

	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/consumer/mailing"
	"github.com/joho/godotenv"
	"github.com/HETIC-MT-P2021/Go_Mail_Groupe05/consumer/rabbit"
)

func main() {
	env, _ := godotenv.Read(".env")

	rabbit.ConnectToRabbit(
		env["RABBIT_HOST"],
		env["RABBIT_PORT"],
		env["RABBIT_USER"],
		env["RABBIT_PASSWORD"])

	smtpPort, _ := strconv.Atoi(env["SMTP_PORT"])

	mailing.InitSMTPCon(env["SMTP_USER"], env["SMTP_PASSWORD"], env["SMTP_HOST"], smtpPort)

	rabbit.Receive()
}
