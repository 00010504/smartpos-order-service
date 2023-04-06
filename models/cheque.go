package models

import "database/sql"

type GetChequeResponse struct {
	Id        string
	CompanyId string
	Name      string
	Message   string
	Logo      *ChequeLogo
	Blocks    []*GetReceiptBlockResponse
}

type ChequeLogo struct {
	Image    string
	ChequeId string
	Left     int8
	Right    int8
	Top      int8
	Bottom   int8
}

type NullChequeLogo struct {
	Image    sql.NullString
	ChequeId sql.NullString
	Left     sql.NullInt16
	Right    sql.NullInt16
	Top      sql.NullInt16
	Bottom   sql.NullInt16
}
