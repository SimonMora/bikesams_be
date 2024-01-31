package models

import (
	"database/sql"
	"time"
)

type Order struct {
	Order_Id       sql.NullInt64
	Order_UserUUID sql.NullString
	Order_AddId    sql.NullInt64
	Order_Date     sql.NullTime
	Order_Total    sql.NullFloat64
	Order_Details  []OrderDetails
}

type OrderDetails struct {
	OD_Id       sql.NullInt64
	OD_OderId   sql.NullInt64
	OD_ProdId   sql.NullInt64
	OD_Quantity sql.NullInt64
	OD_Price    sql.NullFloat64
}

type OrderRequest struct {
	OrderId       int                   `json:"orderId"`
	OrderUserUUID string                `json:"orderUserUUID"`
	OrderAddId    int                   `json:"orderAddId"`
	OrderDate     time.Time             `json:"orderDate"`
	OrderTotal    float64               `json:"orderTotal"`
	OrderDetails  []OrderDetailsRequest `json:"orderDetails"`
}

type OrderDetailsRequest struct {
	ODId       int     `json:"odId"`
	ODOderId   int     `json:"odOrderId"`
	ODProdId   int     `json:"odProdId"`
	ODQuantity int     `json:"odQuantity"`
	ODPrice    float64 `json:"odPrice"`
}

type OrderResponse struct {
	OrderId       int64                  `json:"orderId"`
	OrderUserUUID string                 `json:"orderUserUUID"`
	OrderAddId    int                    `json:"orderAddId"`
	OrderDate     time.Time              `json:"orderDate,omitempty"`
	OrderTotal    float64                `json:"orderTotal"`
	OrderDetails  []OrderDetailsResponse `json:"orderDetails,omitempty"`
}

type OrderDetailsResponse struct {
	ODId       int64   `json:"odId"`
	ODOderId   int     `json:"odOrderId"`
	ODProdId   int     `json:"odProdId"`
	ODQuantity int     `json:"odQuantity"`
	ODPrice    float64 `json:"odPrice"`
}

func (order *OrderResponse) FillOrderReq(req OrderRequest) {
	order.OrderAddId = req.OrderAddId
	order.OrderTotal = req.OrderTotal
	order.OrderUserUUID = req.OrderUserUUID
}

func (order *OrderResponse) FillOrderDb(req Order) {
	order.OrderAddId = int(req.Order_AddId.Int64)
	order.OrderTotal = req.Order_Total.Float64
	order.OrderUserUUID = req.Order_UserUUID.String
	order.OrderId = req.Order_Id.Int64
	order.OrderDate = req.Order_Date.Time
}

func (orderD *OrderDetailsResponse) FillOrderDetailDb(req OrderDetails) {
	orderD.ODId = req.OD_Id.Int64
	orderD.ODOderId = int(req.OD_OderId.Int64)
	orderD.ODPrice = req.OD_Price.Float64
	orderD.ODProdId = int(req.OD_ProdId.Int64)
	orderD.ODQuantity = int(req.OD_Quantity.Int64)
}
