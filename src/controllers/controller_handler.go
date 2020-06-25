package controllers

import "database/sql"

// HandleDbAndSalt is a structure to pass the DB connection and the PW salt as parameters in routes/controllers
type HandleDbAndSalt struct {
	DbCon      *sql.DB
	SaltString string
}

// HandleDb is a structure to pass the DB connection as parameter in routes/controllers
type HandleDb struct {
	DbCon *sql.DB
}
