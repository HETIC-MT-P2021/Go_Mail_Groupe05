package models

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MailingList type
type MailingList struct {
	MailintListID int
	Name          string
	BusinessID    int
}

// MailingListCustomerLink type
type MailingListCustomerLink struct {
	MailintListID int
	CustomerID    int
}

// CreateMailingList will add a new mailingList to the DB
func CreateMailingList(mailingListName string, businessID string) (MailingList, error) {
	var newMailingList MailingList

	createMailingListSQL := `
	INSERT INTO mailing_list (name, business_id)
	VALUES ($1, $2) RETURNING *;`

	mailingRow := db.QueryRow(createMailingListSQL, mailingListName, businessID)
	err := mailingRow.Scan(&newMailingList.MailintListID, &newMailingList.Name, &newMailingList.BusinessID)

	if err != nil {
		return newMailingList, err
	}

	return newMailingList, nil
}

// GetMailingList will get information from a mailingList and it's customers
func GetMailingList(mailingListID string, c *gin.Context) (MailingList, []Customer, error) {
	var thisMailingList MailingList
	var thisCustomers []Customer
	var thisCustomersID []int

	// Get information from this mailingLit
	mailingListSQL := `SELECT * FROM mailing_list WHERE id=$1;`

	mailingRow := db.QueryRow(mailingListSQL, mailingListID)
	err := mailingRow.Scan(&thisMailingList.MailintListID, &thisMailingList.Name, &thisMailingList.BusinessID)

	if err != nil {
		return thisMailingList, thisCustomers, err
	}

	// Get CustomerIDS for this mailing List
	customerLinkSQL := `SELECT * FROM mailing_list_customer_assoc WHERE mailing_list_id=$1;`

	rows, linkCustomerErr := db.QueryContext(c, customerLinkSQL, mailingListID)

	if linkCustomerErr != nil {
		return thisMailingList, thisCustomers, linkCustomerErr
	}

	for rows.Next() {
		var thisCustomerLink MailingListCustomerLink
		if custErr := rows.Scan(&thisCustomerLink.MailintListID, &thisCustomerLink.CustomerID); custErr != nil {
			return thisMailingList, thisCustomers, custErr
		}
		thisCustomersID = append(thisCustomersID, thisCustomerLink.CustomerID)
	}

	if len(thisCustomersID) == 0 {
		return thisMailingList, thisCustomers, errors.New("No customer found for this mailing list")
	}

	// Get information for each customers
	customerSQL := `SELECT * FROM customer WHERE `

	for i := 0; i < len(thisCustomersID); i++ {
		if i != 0 {
			customerSQL += ` OR `
		}
		customerSQL += `id=` + strconv.Itoa(thisCustomersID[i])
	}

	customersRow, customersErr := db.QueryContext(c, customerSQL)

	if customersErr != nil {
		return thisMailingList, thisCustomers, customersErr
	}

	for customersRow.Next() {
		var thisCustomer Customer
		if custErr := customersRow.Scan(&thisCustomer.CustomerID, &thisCustomer.Email, &thisCustomer.Name, &thisCustomer.Surname, &thisCustomer.BusinessID); custErr != nil {
			return thisMailingList, thisCustomers, custErr
		}
		thisCustomers = append(thisCustomers, thisCustomer)
	}

	return thisMailingList, thisCustomers, nil
}
