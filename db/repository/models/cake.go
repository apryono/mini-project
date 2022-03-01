package models

import "time"

// Cake ...
type Cake struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Rating      int        `json:"rating"`
	Image       string     `json:"image"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

var (
	// SelectCakeStatement query statement
	SelectCakeStatement = ` SELECT id, title, description, rating, image, created_at, updated_at from cakes `
)
