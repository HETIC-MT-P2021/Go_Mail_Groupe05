package main

import (
	"strconv"

	dbmodel "packages.hetic.net/gomail/models/db"
	route "packages.hetic.net/gomail/routes"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	env, _ := godotenv.Read(".env")

	dbPort, err := strconv.ParseInt(env["DB_PORT"], 10, 64)

	if err != nil {
		panic(err)
	}

	var dbCon = dbmodel.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], dbPort)

	route.StartRouter(env["API_PORT"], dbCon, env["PW_SALT"])
}
