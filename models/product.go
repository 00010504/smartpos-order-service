package models

import "database/sql"

type NullProduct struct {
	Name     sql.NullString
	Barcode  sql.NullString
	MxikCode sql.NullString
	Image    sql.NullString
	Sku      sql.NullString
}

type NullMeasurementUnit struct {
	Id        sql.NullString
	ShortName sql.NullString
	LongName  sql.NullString
}
