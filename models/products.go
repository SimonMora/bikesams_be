package models

import "time"

type Product struct {
	Prod_Id          int       `json:"prodId"`
	Prod_Title       string    `json:"prodTitle"`
	Prod_Description string    `json:"prodDescription"`
	Prod_CreatedAt   time.Time `json:"prodCreatedAt"`
	Prod_Updated     time.Time `json:"prodUpdated"`
	Prod_Price       float64   `json:"prodPrice"`
	Prod_Path        string    `json:"prodPath"`
	Prod_CategoryId  int       `json:"prodCategoryId"`
	Prod_Stock       int       `json:"prodStock"`
}

type ProductResponse struct {
	Prod_Id          int64     `json:"prodId"`
	Prod_Title       string    `json:"prodTitle"`
	Prod_Description string    `json:"prodDescription"`
	Prod_CreatedAt   time.Time `json:"prodCreatedAt"`
}
