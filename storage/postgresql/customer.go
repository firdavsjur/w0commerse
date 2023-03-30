package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type customerRepo struct {
	db *pgxpool.Pool
}

func NewCustomerRepo(db *pgxpool.Pool) *customerRepo {
	return &customerRepo{
		db: db,
	}
}

func (r *customerRepo) Create(ctx context.Context, req *models.CreateCustomer) (int, error) {
	var (
		query string
		id    int
	)

	// get last id
	query = `
	SELECT
		customer_id
	FROM customers
	ORDER BY customer_id  DESC
	LIMIT 1
`
	err := r.db.QueryRow(ctx, query).Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO customers(
			customer_id, 
			first_name,
			last_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
	)
	VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9)`

	_, err = r.db.Exec(ctx, query,
		id+1,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.Email,
		req.Street,
		req.City,
		req.State,
		req.ZipCode,
	)
	if err != nil {
		return 0, err
	}

	return id + 1, nil
}

func (r *customerRepo) GetByID(ctx context.Context, req *models.CustomerPrimaryKey) (*models.Customer, error) {

	var (
		query    string
		customer models.Customer
	)

	query = `
		SELECT
			customer_id, 
			first_name,
			last_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
		FROM customers
		WHERE customer_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.CustomerId).Scan(
		&customer.CustomerId,
		&customer.FirstName,
		&customer.LastName,
		&customer.Phone,
		&customer.Email,
		&customer.Street,
		&customer.City,
		&customer.State,
		&customer.ZipCode,
	)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepo) GetList(ctx context.Context, req *models.GetListCustomerRequest) (resp *models.GetListCustomerResponse, err error) {

	resp = &models.GetListCustomerResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
				customer_id, 
				first_name,
				last_name,
				phone,
				email,
				street,
				city,
				state,
				zip_code
		FROM customers
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
		var customer models.Customer
		var phone sql.NullString
		err = rows.Scan(
			&resp.Count,
			&customer.CustomerId,
			&customer.FirstName,
			&customer.LastName,
			&phone,
			&customer.Email,
			&customer.Street,
			&customer.City,
			&customer.State,
			&customer.ZipCode,
		)
		if err != nil {
			return nil, err
		}
		customer.Phone = phone.String
		resp.Customers = append(resp.Customers, &customer)

	}

	return resp, nil
}

func (r *customerRepo) Update(ctx context.Context, req *models.UpdateCustomer) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		customers
		SET
			customer_id = :customer_id, 
			customer_name = :customer_name
		WHERE customer_id = :customer_id
	`

	params = map[string]interface{}{
		"customer_id": req.CustomerId,
		"first_name":  req.FirstName,
		"last_name":   req.LastName,
		"phone":       req.Phone,
		"email":       req.Email,
		"street":      req.Street,
		"city":        req.City,
		"state":       req.State,
		"zip_code":    req.ZipCode,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *customerRepo) Delete(ctx context.Context, req *models.CustomerPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM customers
		WHERE customer_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.CustomerId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
