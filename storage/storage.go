package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
	Brand() BrandRepoI
	Product() ProductRepoI
	Stock() StockRepoI
	Customer() CustomerRepoI
	Store() StoreRepoI
	Staff() StaffRepoI
	Order() OrderRepoI
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (int, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Delete(ctx context.Context, req *models.CategoryPrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateCategory) (int64, error)
}
type BrandRepoI interface {
	Create(context.Context, *models.CreateBrand) (int, error)
	GetByID(context.Context, *models.BrandPrimaryKey) (*models.Brand, error)
	GetList(context.Context, *models.GetListBrandRequest) (*models.GetListBrandResponse, error)
	Delete(ctx context.Context, req *models.BrandPrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateBrand) (int64, error)
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (int, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetList(context.Context, *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateProduct) (int64, error)
}

type StockRepoI interface {
	Create(context.Context, *models.Stock) (int, error)
	GetByID(context.Context, *models.GetByIDStockRequest) (*models.GetByIDStockResponse, error)
	GetList(context.Context, *models.GetListStockRequest) (*models.GetListStockResponse, error)
	Delete(ctx context.Context, req *models.GetByIDStockRequest) (int64, error)
	Update(ctx context.Context, req *models.Stock) (int64, error)
}

type CustomerRepoI interface {
	Create(context.Context, *models.CreateCustomer) (int, error)
	GetByID(context.Context, *models.CustomerPrimaryKey) (*models.Customer, error)
	GetList(context.Context, *models.GetListCustomerRequest) (*models.GetListCustomerResponse, error)
	Delete(ctx context.Context, req *models.CustomerPrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateCustomer) (int64, error)
}
type StoreRepoI interface {
	Create(context.Context, *models.CreateStore) (int, error)
	GetByID(context.Context, *models.StorePrimaryKey) (*models.Store, error)
	GetList(context.Context, *models.GetListStoreRequest) (*models.GetListStoreResponse, error)
	Delete(ctx context.Context, req *models.StorePrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateStore) (int64, error)
}

type StaffRepoI interface {
	Create(context.Context, *models.CreateStaff) (int, error)
	GetByID(context.Context, *models.StaffPrimaryKey) (*models.Staff, error)
	GetList(context.Context, *models.GetListStaffRequest) (*models.GetListStaffResponse, error)
	Delete(ctx context.Context, req *models.StaffPrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateStaff) (int64, error)
}
type OrderRepoI interface {
	Create(context.Context, *models.CreateOrder) (int, error)
	GetByID(context.Context, *models.OrderPrimaryKey) (*models.GetByIdOrderResponse, error)
	GetList(context.Context, *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Delete(ctx context.Context, req *models.OrderPrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateOrder) (int64, error)
	CreateOrderItems(ctx context.Context, req *models.CreateOrderItem) (int, error)
	DeleteOrderItem(ctx context.Context, req *models.OrderItemPrimaryKey) (int64, error)
}
