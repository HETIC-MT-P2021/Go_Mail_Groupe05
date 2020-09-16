package models

// Customer type
type Customer struct {
	CustomerID int
	Email      string
	Name       string
	Surname    string
	BusinessID int
}

// CreateAndLinkCustomer will add a new customer and link it to a mailing list in the DB
func CreateAndLinkCustomer(email string, name string, surname string, businessID string, mailingListID string) (Customer, error) {
	var newCustomer Customer

	customerSQL := `
	INSERT INTO customer (email, name, surname, business_id)
	VALUES ($1, $2, $3, $4) RETURNING *;`

	customerRow := db.QueryRow(customerSQL, email, name, surname, businessID)
	custErr := customerRow.Scan(&newCustomer.CustomerID, &newCustomer.Email, &newCustomer.Name, &newCustomer.Surname, &newCustomer.BusinessID)

	if custErr != nil {
		return newCustomer, custErr
	}

	linkSQL := `
	INSERT INTO mailing_list_customer_assoc (mailing_list_id, customer_id)
	VALUES ($1, $2);`

	linkRow := db.QueryRow(linkSQL, mailingListID, newCustomer.CustomerID)
	linkErr := linkRow.Scan()

	if linkErr != nil {
		return newCustomer, linkErr
	}

	return newCustomer, nil
}

// CreateCustomer will add a new customer to the DB
func CreateCustomer(email string, name string, surname string, businessID string) (Customer, error) {
	var newCustomer Customer

	customerSQL := `
	INSERT INTO customer (email, name, surname, business_id)
	VALUES ($1, $2, $3, $4) RETURNING *;`

	customerRow := db.QueryRow(customerSQL, email, name, surname, businessID)
	custErr := customerRow.Scan(&newCustomer.CustomerID, &newCustomer.Email, &newCustomer.Name, &newCustomer.Surname, &newCustomer.BusinessID)

	if custErr != nil {
		return newCustomer, custErr
	}

	return newCustomer, nil
}

// UnlinkCustomerMailingList will remove a customer from a mailing list
func UnlinkCustomerMailingList(customerID string, mailingListID string) error {

	unlinkSQL := `
	DELETE * FROM mailing_list_customer_assoc WHERE mailing_list_id=$1 AND customer_id=$2;`

	customerRow := db.QueryRow(unlinkSQL, mailingListID, customerID)
	unlinkErr := customerRow.Scan()

	if unlinkErr != nil {
		return unlinkErr
	}

	return nil
}

// GetCustomer will get a customer from the DB
func GetCustomer(customerID string) (Customer, error) {
	var thisCustomer Customer

	customerSQL := `
	SELECT * FROM customer WHERE customer_id=$1`

	customerRow := db.QueryRow(customerSQL, customerID)
	custErr := customerRow.Scan(&thisCustomer.CustomerID, &thisCustomer.Email, &thisCustomer.Name, &thisCustomer.Surname, &thisCustomer.BusinessID)

	if custErr != nil {
		return thisCustomer, custErr
	}

	return thisCustomer, nil
}
