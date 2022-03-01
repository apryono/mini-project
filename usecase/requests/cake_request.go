package requests

type CakeRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Rating      int    `json:"rating" validate:"required"`
	Image       string `json:"image"`
}
