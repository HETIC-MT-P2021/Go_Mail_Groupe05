package models

import (
	"database/sql"
	"errors"
)

// Business represent the business type
type Business struct {
	BusinessID int
	Name       string
}

// CreateBusiness will add a new business to the DB
func CreateBusiness(businessName string, db *sql.DB, saltString string) (Business, error) {
	sqlStatement := `
	INSERT INTO business (name)
	VALUES ($1) RETURNING business_id, name;`

	var newBusiness Business

	row := db.QueryRow(sqlStatement, businessName)
	err := row.Scan(&newBusiness.BusinessID, &newBusiness.Name)

	if err != nil {
		return newBusiness, err
	}
	return newBusiness, nil
}

// GetBusiness will return a user from the DB
func GetBusiness(businessID int, db *sql.DB) (Business, error) {

	sqlStatement := `SELECT * FROM business WHERE business_id=$1;`

	var business Business

	row := db.QueryRow(sqlStatement, businessID)

	var err error

	err = row.Scan(&business.BusinessID, &business.Name)

	switch err {
	case sql.ErrNoRows:
		return business, errors.New("Notfound, no business found for this name")
	case nil:
		return business, nil

	default:
		return business, errors.New("Internal Server error")
	}
}
