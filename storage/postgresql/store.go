package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type storeRepo struct {
	db *pgxpool.Pool
}

func NewStoreRepo(db *pgxpool.Pool) *storeRepo {
	return &storeRepo{
		db: db,
	}
}

func (r *storeRepo) Create(ctx context.Context, req *models.CreateStore) (int, error) {
	var (
		query string
		id    int
	)

	// get last id
	query = `
	SELECT
		store_id
	FROM stores
	ORDER BY store_id  DESC
	LIMIT 1
`
	err := r.db.QueryRow(ctx, query).Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO stores(
			store_id, 
			store_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
	)
	VALUES ($1, $2,$3,$4,$5,$6,$7,$8)`

	_, err = r.db.Exec(ctx, query,
		id+1,
		req.StoreName,
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

func (r *storeRepo) GetByID(ctx context.Context, req *models.StorePrimaryKey) (*models.Store, error) {

	var (
		query string
		store models.Store
	)

	query = `
		SELECT
			store_id, 
			store_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
		FROM stores
		WHERE store_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.StoreId).Scan(
		&store.StoreId,
		&store.StoreName,
		&store.Phone,
		&store.Email,
		&store.Street,
		&store.City,
		&store.State,
		&store.ZipCode,
	)
	if err != nil {
		return nil, err
	}

	return &store, nil
}

func (r *storeRepo) GetList(ctx context.Context, req *models.GetListStoreRequest) (resp *models.GetListStoreResponse, err error) {

	resp = &models.GetListStoreResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
				store_id, 
				store_name,
				phone,
				email,
				street,
				city,
				state,
				zip_code
		FROM stores
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
		var store models.Store
		var phone sql.NullString
		err = rows.Scan(
			&resp.Count,
			&store.StoreId,
			&store.StoreName,
			&phone,
			&store.Email,
			&store.Street,
			&store.City,
			&store.State,
			&store.ZipCode,
		)
		if err != nil {
			return nil, err
		}
		store.Phone = phone.String
		resp.Stores = append(resp.Stores, &store)

	}

	return resp, nil
}

func (r *storeRepo) Update(ctx context.Context, req *models.UpdateStore) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		stores
		SET
			store_id = :store_id, 
			store_name = :store_name
		WHERE store_id = :store_id
	`

	params = map[string]interface{}{
		"store_id":   req.StoreId,
		"store_name": req.StoreName,
		"phone":      req.Phone,
		"email":      req.Email,
		"street":     req.Street,
		"city":       req.City,
		"state":      req.State,
		"zip_code":   req.ZipCode,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *storeRepo) Delete(ctx context.Context, req *models.StorePrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM stores
		WHERE store_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.StoreId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
