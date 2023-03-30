package models

type Customer struct {
	CustomerId int    `json:"customerID"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    int    `json:"zipCode"`
}

type CustomerPrimaryKey struct {
	CustomerId int `json:"customerID"`
}

type CreateCustomer struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   int    `json:"zipCode"`
}

type UpdateCustomer struct {
	CustomerId int    `json:"customerID"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    int    `json:"zipCode"`
}

type GetListCustomerRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCustomerResponse struct {
	Count     int         `json:"count"`
	Customers []*Customer `json:"customers"`
}
