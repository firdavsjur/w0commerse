package models

type Store struct {
	StoreId   int    `json:"storeID"`
	StoreName string `json:"storeName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zipCode"`
}

type StorePrimaryKey struct {
	StoreId int `json:"storeID"`
}

type CreateStore struct {
	StoreName string `json:"storeName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zipCode"`
}

type UpdateStore struct {
	StoreId   int    `json:"storeID"`
	StoreName string `json:"storeName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zipCode"`
}

type GetListStoreRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListStoreResponse struct {
	Count  int      `json:"count"`
	Stores []*Store `json:"stores"`
}
