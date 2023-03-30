package models

import "time"

type Order struct {
	OrderId      int       `json:"order_id"`
	CustomerId   int       `json:"customer_id"`
	CustomerData Customer  `json:"customer_data"`
	OrderStatus  int       `json:"order_status"`
	OrderDate    time.Time `json:"order_date"`
	RequiredDate time.Time `json:"required_date"`
	ShippedDate  time.Time `json:"shipped_date"`
	StoreID      int       `json:"store_id"`
	StoreData    Store     `json:"store_data"`
	StaffId      int       `json:"staff_id"`
	StaffData    Staff     `json:"staff_data"`
}
type GetByIdOrderResponse struct {
	OrderId      int          `json:"order_id"`
	CustomerId   int          `json:"customer_id"`
	CustomerData Customer     `json:"customer_data"`
	OrderStatus  int          `json:"order_status"`
	OrderDate    time.Time    `json:"order_date"`
	RequiredDate time.Time    `json:"required_date"`
	ShippedDate  time.Time    `json:"shipped_date"`
	StoreID      int          `json:"store_id"`
	StoreData    Store        `json:"store_data"`
	StaffId      int          `json:"staff_id"`
	StaffData    Staff        `json:"staff_data"`
	Items        []*OrderItem `json:"items"`
}
type OrderPrimaryKey struct {
	OrderId int `json:"order_id"`
}

type CreateOrder struct {
	CustomerId   int       `json:"customer_id"`
	OrderStatus  int       `json:"order_status"`
	OrderDate    time.Time `json:"order_date"`
	RequiredDate time.Time `json:"required_date"`
	ShippedDate  time.Time `json:"shipped_date"`
	StoreID      int       `json:"store_id"`
	StaffId      int       `json:"staff_id"`
}

type UpdateOrder struct {
	OrderId      int       `json:"order_id"`
	CustomerId   int       `json:"customer_id"`
	OrderStatus  int       `json:"order_status"`
	OrderDate    time.Time `json:"order_date"`
	RequiredDate time.Time `json:"required_date"`
	ShippedDate  time.Time `json:"shipped_date"`
	StoreID      int       `json:"store_id"`
	StaffId      int       `json:"staff_id"`
}

type GetListOrderRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count  int      `json:"count"`
	Orders []*Order `json:"orders"`
}

type OrderItem struct {
	OrderId     int     `json:"order_id"`
	ItemId      int     `json:"item_id"`
	ProductId   int     `json:"product_id"`
	ProductData Product `json:"product_data"`
	Quantity    int     `json:"quantity"`
	ListPrice   float32 `json:"list_price"`
	Discount    float32 `json:"discount"`
}

type UpdateOrderItem struct {
	OrderId     int     `json:"order_id"`
	ItemId      int     `json:"item_id"`
	ProductId   int     `json:"product_id"`
	ProductData Product `json:"product_data"`
	Quantity    int     `json:"quantity"`
	ListPrice   float32 `json:"list_price"`
	Discount    float32 `json:"discount"`
}

type CreateOrderItem struct {
	OrderId   int     `json:"order_id"`
	ProductId int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	ListPrice float32 `json:"list_price"`
	Discount  float32 `json:"discount"`
}

type OrderItemPrimaryKey struct {
	OrderId int `json:"order_id"`
	ItemId  int `json:"item_id"`
}
