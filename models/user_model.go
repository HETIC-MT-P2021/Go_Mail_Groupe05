package models

import (
	"database/sql"
	"errors"

	"packages.hetic.net/gomail/utils"
)

// UnSavedUser represent the user's type before they are saved in the DB
type UnSavedUser struct {
	UserID     int
	Email      string
	Password   string
	BusinessID int
}

// SavedUser represent the user's type after they are saved in the DB
type SavedUser struct {
	UserID     int
	Email      string
	Password   []byte
	BusinessID int
}

// GetUser will return a user from the DB
func GetUser(email string, getPassword bool) (SavedUser, error) {
	var sqlStatement string

	if getPassword == true {
		sqlStatement = `SELECT * FROM users WHERE email=$1;`
	} else {
		sqlStatement = `SELECT id, email, business_id FROM users WHERE email=$1;`
	}

	var user SavedUser

	row := db.QueryRow(sqlStatement, email)
	var err error

	if getPassword == true {
		err = row.Scan(&user.UserID, &user.Email, &user.Password, &user.BusinessID)
	} else {

		err = row.Scan(&user.UserID, &user.Email, &user.BusinessID)
	}

	switch err {
	case sql.ErrNoRows:
		return user, errors.New("notfound, no user found for this email")
	case nil:
		return user, nil

	default:
		return user, errors.New("internal Server error")
	}
}

// CreateUser will add a new user to the DB
func CreateUser(email string, password string, businessID string) (SavedUser, error) {
	hashedPassword := utils.HashPassword(password)

	sqlStatement := `
	INSERT INTO users (email, password, business_id)
	VALUES ($1, $2, $3) RETURNING id, email, business_id;`

	var newUser SavedUser

	row := db.QueryRow(sqlStatement, email, hashedPassword, businessID)
	err := row.Scan(&newUser.UserID, &newUser.Email, &newUser.BusinessID)

	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// VerifyUserCredentials will fetch user's password and compare it with the one entered
// Recommanded to use at program's end
func VerifyUserCredentials(email string, password string) (bool, error) {
	thisUser, err := GetUser(email, true)

	if err != nil {
		return false, err
	}

	return utils.CheckPass(password, thisUser.Password), nil
}
