package models

import "database/sql"

type ShortShop struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	CompanyId string `json:"company_id"`
}

type NullShop struct {
	Id        sql.NullString
	Title     sql.NullString
	CompanyId sql.NullString
}
