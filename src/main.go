package main

import "gopkg.in/gomail.v2"

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "SENDER_EMAIL")
	m.SetHeader("To", "DEST_1_EMAIL", "DEST_2_EMAIL")
	m.SetHeader("Subject", "SUBJECT")
	m.SetBody("text/plain", "MESSAGE_IN_STRING")
	m.SetBody("text/html", "MESSAGE_IN_HTML")

	d := gomail.NewDialer("smtp.gmail.com", 465, "SENDER_EMAIL", "SENDER_GMAIL_PASSWORD")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	// env, _ := godotenv.Read(".env")

	// dbPort, err := strconv.ParseInt(env["DB_PORT"], 10, 64)

	// if err != nil {
	// 	panic(err)
	// }

	// var dbCon = model.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], dbPort)

	// route.StartRouter(env["API_PORT"], dbCon, env["PW_SALT"])
}
