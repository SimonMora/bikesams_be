package models

import "time"

type Product struct {
	ProdId          int64     `json:"prodId"`
	ProdTitle       string    `json:"prodTitle"`
	ProdDescription string    `json:"prodDescription"`
	ProdCreatedAt   time.Time `json:"prodCreatedAt"`
	ProdUpdated     time.Time `json:"prodUpdated"`
	ProdPrice       float64   `json:"prodPrice,omitempty"`
	ProdPath        string    `json:"prodPath"`
	ProdCategoryId  int       `json:"prodCategoryId"`
	ProdStock       int       `json:"prodStock"`
}

type ProductResponse struct {
	ProdId          int64     `json:"prodId"`
	ProdTitle       string    `json:"prodTitle"`
	ProdDescription string    `json:"prodDescription"`
	ProdCreatedAt   time.Time `json:"prodCreatedAt"`
}

type ProductRequest struct {
	ProdId          int64     `json:"prodId"`
	ProdTitle       string    `json:"prodTitle"`
	ProdDescription string    `json:"prodDescription"`
	ProdCreatedAt   time.Time `json:"prodCreatedAt"`
	ProdUpdated     time.Time `json:"prodUpdated"`
	ProdPrice       float64   `json:"prodPrice,omitempty"`
	ProdPath        string    `json:"prodPath"`
	ProdCategoryId  int       `json:"prodCategoryId"`
	ProdStock       int       `json:"prodStock"`
	ProdSearch      string    `json:"search,omitempty"`
	ProdCategPath   string    `json:"categPath,omitempty"`
}
