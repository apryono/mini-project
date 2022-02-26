package bootstrap

import (
	"github.com/mini-project/usecase"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"

	"github.com/gofiber/fiber/v2"
)

//Bootstrap ...
type Bootstrap struct {
	App        *fiber.App
	ContractUC usecase.ContractUC
	Validator  *validator.Validate
	Translator ut.Translator
}
