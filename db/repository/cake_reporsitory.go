package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/mini-project/db/repository/models"
	"github.com/mini-project/pkg/str"
)

type ICakeRepository interface {
	Add(c context.Context, model *models.Cake) (int, error)
	FindByID(c context.Context, id int) (models.Cake, error)
	FindAllCake(c context.Context, param models.CakeParameter) ([]models.Cake, error)
	Edit(c context.Context, req models.Cake) (int, error)
	Delete(c context.Context, id int) error
}

type CakeRepository struct {
	DB *sql.DB
	Tx *sql.Tx
}

func NewCakeRepository(DB *sql.DB, Tx *sql.Tx) ICakeRepository {
	return &CakeRepository{
		DB: DB,
		Tx: Tx,
	}
}

// Add cake to database
func (r *CakeRepository) Add(c context.Context, model *models.Cake) (res int, err error) {
	statement := `INSERT INTO cakes (title, description, rating, image, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6) returning id`

	if r.Tx != nil {
		err = r.Tx.QueryRowContext(c, statement, model.Title, model.Description, model.Rating,
			model.Image, model.CreatedAt, model.UpdatedAt).Scan(&res)
	} else {
		err = r.DB.QueryRowContext(c, statement, model.Title, model.Description, model.Rating,
			model.Image, model.CreatedAt, model.UpdatedAt).Scan(&res)
	}

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *CakeRepository) FindByID(c context.Context, id int) (res models.Cake, err error) {
	statement := str.Spacing(models.SelectCakeStatement, ` WHERE id = $1`)

	row := r.DB.QueryRowContext(c, statement, id)
	res, err = r.scanRow(row)
	if err != nil {
		return res, err
	}

	return res, err

}

func (r *CakeRepository) scanRow(row *sql.Row) (res models.Cake, err error) {
	err = row.Scan(
		&res.ID, &res.Title, &res.Description, &res.Rating, &res.Image, &res.CreatedAt, &res.UpdatedAt,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *CakeRepository) scanRows(rows *sql.Rows) (res models.Cake, err error) {
	err = rows.Scan(
		&res.ID, &res.Title, &res.Description, &res.Rating, &res.Image, &res.CreatedAt, &res.UpdatedAt,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *CakeRepository) FindAllCake(c context.Context, param models.CakeParameter) (res []models.Cake, err error) {
	var conditionString string
	if param.Search != "" {
		search := strings.ToLower(param.Search)
		conditionString += ` WHERE lower(title) LIKE '%` + search + `%' OR lower(description) LIKE '%` + search + `%'`
	}

	statement := str.Spacing(models.SelectCakeStatement, conditionString)

	rows, err := r.DB.QueryContext(c, statement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		temp, err := r.scanRows(rows)
		if err != nil {
			return res, err
		}

		res = append(res, temp)
	}

	return res, err
}

func (r *CakeRepository) Edit(c context.Context, req models.Cake) (res int, err error) {
	date := time.Now()
	statement := ` UPDATE cakes SET title = $1, description = $2, rating = $3, 
		image = $4, updated_at = $5 WHERE id = $6 returning id`

	if r.Tx != nil {
		err = r.Tx.QueryRowContext(c, statement,
			req.Title, req.Description, req.Rating, req.Image, date, req.ID).Scan(&res)
	} else {
		r.DB.QueryRowContext(c, statement,
			req.Title, req.Description, req.Rating, req.Image, date, req.ID).Scan(&res)
	}

	if err != nil {
		return res, err
	}

	return res, err
}

func (r *CakeRepository) Delete(c context.Context, id int) (err error) {
	statement := ` DELETE FROM cakes where id = $1`

	if r.Tx != nil {
		_, err = r.Tx.ExecContext(c, statement, id)
	} else {
		_, err = r.DB.ExecContext(c, statement, id)
	}

	if err != nil {
		return err
	}

	return err
}
