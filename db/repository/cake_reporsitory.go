package repository

import (
	"context"
	"database/sql"

	"github.com/mini-project/db/repository/models"
)

type ICakeRepository interface {
	Add(c context.Context, model *models.Cake) (int, error)
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
