package models

import "database/sql"

type ShortCompany struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type NullCompany struct {
	Id   sql.NullString
	Name sql.NullString
}
