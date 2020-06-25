package models

import (
	"database/sql"
	"fmt"

	// Import handle for postgres
	_ "github.com/lib/pq"
)

// ConnectToDB Set up connection to the postgres DB
// Will panic on error
func ConnectToDB(host string, dbname string, user string, password string, port int64) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Database connection params error")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Database initialisation error")
		panic(err)
	}

	fmt.Println("Database successfully connected!")

	return db
}

// CloseDbConnection will end dialogue with the DB
// Recommanded to use at program's end
func CloseDbConnection(db *sql.DB) {
	defer db.Close()
	fmt.Println("DB is closed")
}
