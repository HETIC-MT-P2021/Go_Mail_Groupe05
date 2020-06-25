package models

import "database/sql"

// Campaign Represents a mail campaign
type Campaign struct {
	CampaignID    int
	Name          string
	MailingListID int
	BusinessID    int
}

// CreateCampaignWithExistingMailingList will add a new Campaign to the DB with an existing mailingList
func CreateCampaignWithExistingMailingList(campaignName string, mailingListID int, businessID int, db *sql.DB) (Campaign, error) {
	sqlStatement := `
	INSERT INTO campaign (name, mailing_list_id, business_id)
	VALUES ($1, $2, $3) RETURNING *;`

	var newCampaign Campaign

	row := db.QueryRow(sqlStatement, campaignName, mailingListID, businessID)
	err := row.Scan(&newCampaign.CampaignID, &newCampaign.Name, &newCampaign.MailingListID, &newCampaign.BusinessID)

	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}

// CreateCampaignAndMailingList will add a new Campaign and a new mailingList to the DB
func CreateCampaignAndMailingList(mailingListName string, campaignName string, businessID int, db *sql.DB) (MailingList, Campaign, error) {
	var newMailingList MailingList
	var newCampaign Campaign

	createMailingListSQL := `
	INSERT INTO mailing_list (name, business_id)
	VALUES ($1, $2) RETURNING *;`

	mailingRow := db.QueryRow(createMailingListSQL, mailingListName, businessID)
	err := mailingRow.Scan(&newMailingList.MailintListID, &newMailingList.Name, &newMailingList.BusinessID)

	if err != nil {
		return newMailingList, newCampaign, err
	}

	createCampaignSQL := `
	INSERT INTO campaign (name, mailing_list_id, business_id)
	VALUES ($1, $2, $3) RETURNING *;`

	campaignRow := db.QueryRow(createCampaignSQL, campaignName, newMailingList.MailintListID, businessID)
	err = campaignRow.Scan(&newCampaign.CampaignID, &newCampaign.Name, &newCampaign.MailingListID, &newCampaign.BusinessID)

	if err != nil {
		return newMailingList, newCampaign, err
	}
	return newMailingList, newCampaign, nil
}
