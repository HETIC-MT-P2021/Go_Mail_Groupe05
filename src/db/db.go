package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"packages.hetic.net/gomail/hash"
)

// UnSavedUser represent the user's type before they are saved in the DB
type UnSavedUser struct {
	UserID       int
	Email        string
	Password     string
	EnterpriseID int
}

// SavedUser represent the user's type after they are saved in the DB
type SavedUser struct {
	UserID       int
	Email        string
	Password     []byte
	EnterpriseID int
}

// ConnectToDB Set up connection to the postgres DB
// Will panic on error
func ConnectToDB(host string, dbname string, user string, password string, port int64) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database successfully connected!")

	return db
}

// GetUser will return a user from the DB
// Will panic on error
func GetUser(email string, db *sql.DB, getPassword bool) (SavedUser, error) {
	var sqlStatement string

	if getPassword == true {
		sqlStatement = `SELECT * FROM users WHERE email=$1;`
	} else {
		sqlStatement = `SELECT user_id, email, enterprise_id FROM users WHERE email=$1;`
	}

	var user SavedUser

	row := db.QueryRow(sqlStatement, email)
	var err error

	if getPassword == true {
		err = row.Scan(&user.UserID, &user.Email, &user.Password, &user.EnterpriseID)
	} else {

		err = row.Scan(&user.UserID, &user.Email, &user.EnterpriseID)
	}

	switch err {
	case sql.ErrNoRows:
		fmt.Println("User does not exist!")
		return user, nil
	case nil:
		return user, nil

	default:
		panic(err)
	}
}

// CreateUser will add a new user to the DB
// Will panic on error
func CreateUser(user UnSavedUser, db *sql.DB, saltString string) (SavedUser, error) {
	hashedPassword := hash.HashPassword(user.Password, saltString)
	fmt.Println(hashedPassword)
	sqlStatement := `
	INSERT INTO users (email, password, enterprise_id)
	VALUES ($1, $2, $3) RETURNING user_id, email, enterprise_id;`

	var newUser SavedUser

	row := db.QueryRow(sqlStatement, user.Email, hashedPassword, user.EnterpriseID)
	err := row.Scan(&newUser.UserID, &newUser.Email, &newUser.EnterpriseID)

	if err != nil {
		panic(err)
	}
	fmt.Println("Utiliateur Créé !")
	return newUser, err
}

// CloseDbConnection will end dialogue with the DB
// Recommanded to use at program's end
func CloseDbConnection(db *sql.DB) {
	defer db.Close()
	fmt.Println("DB is closed")
}

// VerifyUserCredentials will fetch user's password and compare it with the one entered
// Recommanded to use at program's end
func VerifyUserCredentials(email string, password string, dbConnection *sql.DB, saltString string) bool {
	thisUser, err := GetUser(email, dbConnection, true)

	if err != nil {
		return false
	}

	return hash.CheckPass(password, thisUser.Password, saltString)
}
