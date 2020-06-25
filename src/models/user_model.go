package models

import (
	"database/sql"
	"errors"

	"packages.hetic.net/gomail/utils"
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

// GetUser will return a user from the DB
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
		return user, errors.New("Notfound, no user found for this email")
	case nil:
		return user, nil

	default:
		return user, errors.New("Internal Server error")
	}
}

// CreateUser will add a new user to the DB
func CreateUser(user UnSavedUser, db *sql.DB, saltString string) (SavedUser, error) {
	hashedPassword := utils.HashPassword(user.Password, saltString)

	sqlStatement := `
	INSERT INTO users (email, password, enterprise_id)
	VALUES ($1, $2, $3) RETURNING user_id, email, enterprise_id;`

	var newUser SavedUser

	row := db.QueryRow(sqlStatement, user.Email, hashedPassword, user.EnterpriseID)
	err := row.Scan(&newUser.UserID, &newUser.Email, &newUser.EnterpriseID)

	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// VerifyUserCredentials will fetch user's password and compare it with the one entered
// Recommanded to use at program's end
func VerifyUserCredentials(email string, password string, dbConnection *sql.DB, saltString string) (bool, error) {
	thisUser, err := GetUser(email, dbConnection, true)

	if err != nil {
		return false, err
	}

	return utils.CheckPass(password, thisUser.Password, saltString), nil
}
