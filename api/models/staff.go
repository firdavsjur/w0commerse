package models

type Staff struct {
	StaffId   int    `json:"staffID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Active    int    `json:"active"`
	StoreId   int    `json:"storeID"`
	ManagerId int    `json:"managerID"`
}

type StaffPrimaryKey struct {
	StaffId int `json:"staffID"`
}

type CreateStaff struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Active    int    `json:"active"`
	StoreId   int    `json:"storeID"`
	ManagerId int    `json:"managerID"`
}

type UpdateStaff struct {
	StaffId   int    `json:"staffID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Active    int    `json:"active"`
	StoreId   int    `json:"storeID"`
	ManagerId int    `json:"managerID"`
}

type GetListStaffRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListStaffResponse struct {
	Count  int      `json:"count"`
	Staffs []*Staff `json:"staffs"`
}
