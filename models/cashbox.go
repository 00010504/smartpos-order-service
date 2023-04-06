package models

import "database/sql"

type Cashbox struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	ShopId    string `json:"shop_id"`
	CompanyId string `json:"company_id"`
}

type NullCashbox struct {
	Id        sql.NullString
	Title     sql.NullString
	ShopId    sql.NullString
	CompanyId sql.NullString
}

type GetCashboxResponse struct {
	Id      string
	Title   string
	Shop    *ShortShop
	Company *ShortCompany
}
