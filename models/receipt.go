package models

type GetReceiptBlockResponse struct {
	Id             string
	Name           string
	NameTr         map[string]string
	Fields         []*GetReceiptFieldResponse
}

type GetReceiptFieldResponse struct {
	Id     string
	Name   string
	NameTr map[string]string
}
