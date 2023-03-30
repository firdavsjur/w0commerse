package models

type Stock struct {
	StoreID   int `json:"storeID"`
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

type GetByIDStockRequest struct {
	StoreID int `json:"storeID"`
}

type GetByIDStockResponse struct {
	Count  int `json:"count"`
	Stocks []*Stock
}

type GetListStockRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListStockResponse struct {
	Count  int `json:"count"`
	Stocks []*Stock
}
