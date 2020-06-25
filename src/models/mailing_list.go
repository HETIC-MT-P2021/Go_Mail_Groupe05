package models

import "database/sql"

// MailingList type
type MailingList struct {
	MailintListID int
	Name          string
	BusinessID    int
}

// CreateMailingList will add a new mailingList to the DB
func CreateMailingList(mailingListName string, businessID int, db *sql.DB) (MailingList, error) {
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
