package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) *staffRepo {
	return &staffRepo{
		db: db,
	}
}

func (r *staffRepo) Create(ctx context.Context, req *models.CreateStaff) (int, error) {
	var (
		query string
		id    int
	)

	// get last id
	query = `
	SELECT
		staff_id
	FROM staffs
	ORDER BY staff_id  DESC
	LIMIT 1
`
	err := r.db.QueryRow(ctx, query).Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO staffs(
			staff_id, 
			first_name,
			last_name,
			phone,
			email,
			active,
			store_id,
			manager_id
	)
	VALUES ($1, $2,$3,$4,$5,$6,$7,$8)`

	_, err = r.db.Exec(ctx, query,
		id+1,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.Email,
		req.Active,
		req.StoreId,
		req.ManagerId,
	)
	if err != nil {
		return 0, err
	}

	return id + 1, nil
}

func (r *staffRepo) GetByID(ctx context.Context, req *models.StaffPrimaryKey) (*models.Staff, error) {

	var (
		query string
		staff models.Staff
	)

	query = `
		SELECT
			staff_id, 
			first_name,
			last_name,
			phone,
			email,
			active,
			store_id,
			manager_id
		FROM staffs
		WHERE staff_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.StaffId).Scan(
		&staff.StaffId,
		&staff.FirstName,
		&staff.LastName,
		&staff.Phone,
		&staff.Email,
		&staff.Active,
		&staff.StoreId,
		&staff.ManagerId,
	)
	if err != nil {
		return nil, err
	}

	return &staff, nil
}

func (r *staffRepo) GetList(ctx context.Context, req *models.GetListStaffRequest) (resp *models.GetListStaffResponse, err error) {

	resp = &models.GetListStaffResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
				staff_id, 
				first_name,
				last_name,
				phone,
				email,
				active,
				store_id,
				manager_id
		FROM staffs
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
		var staff models.Staff
		var phone sql.NullString
		err = rows.Scan(
			&resp.Count,
			&staff.StaffId,
			&staff.FirstName,
			&staff.LastName,
			&phone,
			&staff.Email,
			&staff.Active,
			&staff.StoreId,
			&staff.ManagerId,
		)
		if err != nil {
			return nil, err
		}
		staff.Phone = phone.String
		resp.Staffs = append(resp.Staffs, &staff)

	}

	return resp, nil
}

func (r *staffRepo) Update(ctx context.Context, req *models.UpdateStaff) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		staffs
		SET
			staff_id = :staff_id, 
			staff_name = :staff_name
		WHERE staff_id = :staff_id
	`

	params = map[string]interface{}{
		"staff_id":   req.StaffId,
		"first_name": req.FirstName,
		"last_name":  req.LastName,
		"phone":      req.Phone,
		"email":      req.Email,
		"active":     req.Active,
		"store_id":   req.StoreId,
		"manager_id": req.ManagerId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *staffRepo) Delete(ctx context.Context, req *models.StaffPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM staffs
		WHERE staff_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.StaffId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
