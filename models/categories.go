package models

type Category struct {
	Categ_Id   int    `json:"categId"`
	Categ_Name string `json:"categName"`
	Categ_Path string `json:"categPath"`
}

type CategoryProcessResult struct {
	CategId int64 `json:"categId"`
}
