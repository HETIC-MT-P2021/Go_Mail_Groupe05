package consumer

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
	"packages.hetic.net/gomail/consumer/mailing"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

func main() {
	env, _ := godotenv.Read(".env")

	smtpPort, _ := strconv.Atoi(env["SMTP_PORT"])

	mailing.InitSMTPCon(env["SMTP_USER"], env["SMTP_PASSWORD"], env["SMTP_HOST"], smtpPort)
}
