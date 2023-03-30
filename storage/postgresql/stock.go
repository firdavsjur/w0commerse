package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type stockRepo struct {
	db *pgxpool.Pool
}

func NewStockRepo(db *pgxpool.Pool) *stockRepo {
	return &stockRepo{
		db: db,
	}
}

func (r *stockRepo) Create(ctx context.Context, req *models.Stock) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO stocks(
			store_id, 
			product_id,
			quantity

	)
	VALUES ($1, $2,$3)`

	_, err := r.db.Exec(ctx, query,
		req.StoreID,
		req.ProductID,
		req.Quantity,
	)
	if err != nil {
		return 0, err
	}

	return id + 1, nil
}

func (r *stockRepo) GetByID(ctx context.Context, req *models.GetByIDStockRequest) (*models.GetByIDStockResponse, error) {

	var (
		query string
		resp  *models.GetByIDStockResponse
	)
	resp = &models.GetByIDStockResponse{}

	query = `
		SELECT
		COUNT(*) OVER(),
		store_id,
		product_id,
		quantity
	FROM stocks
		WHERE store_id = $1
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(
			&resp.Count,
			&stock.StoreID,
			&stock.ProductID,
			&stock.Quantity,
		)
		if err != nil {
			return nil, err
		}

		resp.Stocks = append(resp.Stocks, &stock)
	}

	return resp, nil
}

func (r *stockRepo) GetList(ctx context.Context, req *models.GetListStockRequest) (resp *models.GetListStockResponse, err error) {

	resp = &models.GetListStockResponse{}

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
			product_id,
			quantity
		FROM stocks
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
		var stock models.Stock
		err = rows.Scan(
			&resp.Count,
			&stock.StoreID,
			&stock.ProductID,
			&stock.Quantity,
		)
		if err != nil {
			return nil, err
		}

		resp.Stocks = append(resp.Stocks, &stock)
	}

	return resp, nil
}

func (r *stockRepo) Update(ctx context.Context, req *models.Stock) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		stocks
		SET
			stock_id = :stock_id, 
			stock_name = :stock_name
		WHERE stock_id = :stock_id
	`

	params = map[string]interface{}{
		"store_id":   req.StoreID,
		"product_id": req.ProductID,
		"quantity":   req.Quantity,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *stockRepo) Delete(ctx context.Context, req *models.GetByIDStockRequest) (int64, error) {
	query := `
		DELETE 
		FROM stocks
		WHERE store_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.StoreID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
