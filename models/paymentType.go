package models

import "database/sql"

type PaymentType struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
}

type NullPaymentType struct {
	Id        sql.NullString
	Name      sql.NullString
	CompanyId sql.NullString
}

type PaymentsType struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	ShopId   string  `json:"shop_id"`
	ShopName string  `json:"shop_name"`
}
