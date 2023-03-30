package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) Create(ctx context.Context, req *models.CreateOrder) (int, error) {
	var (
		query string
		id    int
	)

	// get last id
	query = `
	SELECT
		order_id
	FROM orders
	ORDER BY order_id  DESC
	LIMIT 1
`
	err := r.db.QueryRow(ctx, query).Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO orders(
			order_id, 
			customer_id,
			order_status,
			order_date,
			required_date,
			shipped_date,
			store_id,
			staff_id
	)
	VALUES ($1, $2,$3,$4,$5,$6,$7,$8)`

	_, err = r.db.Exec(ctx, query,
		id+1,
		req.CustomerId,
		req.OrderStatus,
		req.OrderDate,
		req.RequiredDate,
		req.ShippedDate,
		req.StoreID,
		req.StaffId,
	)
	if err != nil {
		return 0, err
	}

	return id + 1, nil
}

func (r *orderRepo) GetByID(ctx context.Context, req *models.OrderPrimaryKey) (*models.GetByIdOrderResponse, error) {

	var (
		query         string
		order         models.GetByIdOrderResponse
		customerPhone sql.NullString
		// pg_order_id
		// pg_item_id
		// pg_product_id
		// pg_quantity
		// pg_list_price
		// pg_discount
	)

	query = `
		SELECT
			ord.order_id, 
			ord.customer_id,
			
			cu.customer_id,
			cu.first_name,
			cu.last_name,
			cu.phone,
			cu.email,
			cu.street,
			cu.city,
			cu.state,
			cu.zip_code,

			ord.order_status,
			ord.order_date,
			ord.required_date,
			ord.shipped_date,
			ord.store_id,
			
			st.store_id,
			st.store_name,
			st.phone,
			st.email,
			st.street,
			st.city,
			st.state,
			st.zip_code,

			ord.staff_id,
			sta.staff_id,
			sta.first_name,
			sta.last_name,
			sta.phone,
			sta.email,
			sta.active,
			sta.store_id,
			sta.manager_id
			

		FROM orders as ord
		JOIN customers as cu on cu.customer_id = ord.customer_id
		JOIN stores as st on st.store_id = ord.store_id
		JOIN staffs as sta on sta.staff_id  = ord.staff_id
		WHERE ord.order_id = $1
		`

	err := r.db.QueryRow(ctx, query, req.OrderId).Scan(
		&order.OrderId,
		&order.CustomerId,
		&order.CustomerData.CustomerId,
		&order.CustomerData.FirstName,
		&order.CustomerData.LastName,
		&customerPhone,
		&order.CustomerData.Email,
		&order.CustomerData.Street,
		&order.CustomerData.City,
		&order.CustomerData.State,
		&order.CustomerData.ZipCode,
		&order.OrderStatus,
		&order.OrderDate,
		&order.RequiredDate,
		&order.ShippedDate,
		&order.StoreID,
		&order.StoreData.StoreId,
		&order.StoreData.StoreName,
		&order.StoreData.Phone,
		&order.StoreData.Email,
		&order.StoreData.Street,
		&order.StoreData.City,
		&order.StoreData.State,
		&order.StoreData.ZipCode,
		&order.StaffId,
		&order.StaffData.StaffId,
		&order.StaffData.FirstName,
		&order.StaffData.LastName,
		&order.StaffData.Phone,
		&order.StaffData.Email,
		&order.StaffData.Active,
		&order.StaffData.StoreId,
		&order.StaffData.ManagerId,
	)
	order.CustomerData.Phone = customerPhone.String
	if err != nil {
		return nil, err
	}

	query = `
		SELECT
		 ori.order_id,
		 ori.item_id,
		 ori.product_id,
		 
		pr.product_id,
		pr.product_name,
		pr.brand_id,
		pr.category_id,
		pr.model_year,
		pr.list_price,

		ori.quantity,
		ori.list_price,
		ori.discount
		from order_items as ori
		JOIN products as pr on pr.product_id = ori.product_id
		
		where order_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.OrderId)
	if err != nil {
		return nil, fmt.Errorf("sdas")
	}
	defer rows.Close()

	for rows.Next() {
		var orderItem models.OrderItem
		err = rows.Scan(
			&orderItem.OrderId,
			&orderItem.ItemId,
			&orderItem.ProductId,
			&orderItem.ProductData.ProductId,
			&orderItem.ProductData.ProductName,
			&orderItem.ProductData.BrandId,
			&orderItem.ProductData.CategoryId,
			&orderItem.ProductData.ModelYear,
			&orderItem.ProductData.ListPrice,
			&orderItem.Quantity,
			&orderItem.ListPrice,
			&orderItem.Discount,
		)
		if err != nil {
			return nil, err
		}
		order.Items = append(order.Items, &orderItem)

	}

	return &order, nil
}

func (r *orderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error) {

	resp = &models.GetListOrderResponse{}

	var (
		query         string
		filter        = " WHERE TRUE "
		offset        = " OFFSET 0"
		limit         = " LIMIT 10"
		customerPhone sql.NullString
	)

	query = `
		
			
				SELECT
				COUNT(*) OVER(),
				ord.order_id, 
				ord.customer_id,
				
				cu.customer_id,
				cu.first_name,
				cu.last_name,
				cu.phone,
				cu.email,
				cu.street,
				cu.city,
				cu.state,
				cu.zip_code,

				ord.order_status,
				ord.order_date,
				ord.required_date,
				ord.shipped_date,
				ord.store_id,
				
				st.store_id,
				st.store_name,
				st.phone,
				st.email,
				st.street,
				st.city,
				st.state,
				st.zip_code,

				ord.staff_id,
				sta.staff_id,
				sta.first_name,
				sta.last_name,
				sta.phone,
				sta.email,
				sta.active,
				sta.store_id,
				sta.manager_id

			FROM orders as ord
			JOIN customers as cu on cu.customer_id = ord.customer_id
			JOIN stores as st on st.store_id = ord.store_id
			JOIN staffs as sta on sta.staff_id  = ord.staff_id
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&resp.Count,
			&order.OrderId,
			&order.CustomerId,
			&order.CustomerData.CustomerId,
			&order.CustomerData.FirstName,
			&order.CustomerData.LastName,
			&customerPhone,
			&order.CustomerData.Email,
			&order.CustomerData.Street,
			&order.CustomerData.City,
			&order.CustomerData.State,
			&order.CustomerData.ZipCode,
			&order.OrderStatus,
			&order.OrderDate,
			&order.RequiredDate,
			&order.ShippedDate,
			&order.StoreID,
			&order.StoreData.StoreId,
			&order.StoreData.StoreName,
			&order.StoreData.Phone,
			&order.StoreData.Email,
			&order.StoreData.Street,
			&order.StoreData.City,
			&order.StoreData.State,
			&order.StoreData.ZipCode,
			&order.StaffId,
			&order.StaffData.StaffId,
			&order.StaffData.FirstName,
			&order.StaffData.LastName,
			&order.StaffData.Phone,
			&order.StaffData.Email,
			&order.StaffData.Active,
			&order.StaffData.StoreId,
			&order.StaffData.ManagerId,
		)
		order.CustomerData.Phone = customerPhone.String
		if err != nil {
			return nil, err
		}
		resp.Orders = append(resp.Orders, &order)

	}

	return resp, nil
}

func (r *orderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		orders
		SET
			order_id = :order_id, 
			order_name = :order_name
		WHERE order_id = :order_id
	`

	params = map[string]interface{}{
		"order_id":      req.OrderId,
		"customer_id":   req.CustomerId,
		"order_status":  req.OrderStatus,
		"order_date":    req.OrderDate,
		"required_date": req.RequiredDate,
		"shipped_date":  req.ShippedDate,
		"store_id":      req.StoreID,
		"staff_id":      req.StaffId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *orderRepo) Delete(ctx context.Context, req *models.OrderPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM orders
		WHERE order_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.OrderId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *orderRepo) CreateOrderItems(ctx context.Context, req *models.CreateOrderItem) (int, error) {
	var (
		query string
		id    int
	)

	// get last id
	query = `
	SELECT
		item_id
	FROM order_items
	ORDER BY item_id  DESC
	LIMIT 1
`
	err := r.db.QueryRow(ctx, query).Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO order_items(
			order_id, 
			item_id,
			product_id,
			quantity,
			list_price,
			discount
	)
	VALUES ($1, $2,$3,$4,$5,$6)`

	_, err = r.db.Exec(ctx, query,

		req.OrderId,
		id+1,
		req.ProductId,
		req.Quantity,
		req.ListPrice,
		req.Discount,
	)
	if err != nil {
		return 0, err
	}

	return req.OrderId, nil
}

func (r *orderRepo) DeleteOrderItem(ctx context.Context, req *models.OrderItemPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM order_items
		WHERE item_id = $1 and order_id = $2
	`

	result, err := r.db.Exec(ctx, query, req.ItemId, req.OrderId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
