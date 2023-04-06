package models

import "database/sql"

type Client struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	CompanyId   string `json:"company_id"`
}

type NullClient struct {
	Id          sql.NullString
	FirstName   sql.NullString
	LastName    sql.NullString
	PhoneNumber sql.NullString
	CompanyId   sql.NullString
}
