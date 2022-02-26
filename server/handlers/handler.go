package handlers

import (
	"database/sql"

	"github.com/mini-project/usecase"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	FiberApp   *fiber.App
	ContractUC *usecase.ContractUC
	Db         *sql.DB
	Validator  *validator.Validate
	Translator ut.Translator
}
