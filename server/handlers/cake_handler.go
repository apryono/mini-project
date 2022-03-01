package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mini-project/db/repository/models"
	"github.com/mini-project/pkg/str"
	"github.com/mini-project/usecase"
	"github.com/mini-project/usecase/requests"
)

type CakeHandler struct {
	Handler
}

func (h *Handler) Add(ctx *fiber.Ctx) error {
	c := ctx.Context()

	input := new(requests.CakeRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.CakeUC{ContractUC: h.ContractUC}
	res, err := uc.AddCake(c, input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err.Error(), http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, nil, 0)
}

func (h *Handler) DetailCake(ctx *fiber.Ctx) error {
	c := ctx.Context()
	reqID := ctx.Params("id")
	ID := str.StringToInt(reqID)

	uc := usecase.CakeUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, ID)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err.Error(), http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, nil, 0)
}

func (h *Handler) ListAllCake(ctx *fiber.Ctx) error {
	c := ctx.Context()

	param := models.CakeParameter{
		Search: ctx.Query("search"),
	}

	uc := usecase.CakeUC{ContractUC: h.ContractUC}
	res, err := uc.FindAllCake(c, param)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err.Error(), http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, nil, 0)
}

func (h *Handler) Edit(ctx *fiber.Ctx) error {
	c := ctx.Context()
	reqID := ctx.Params("id")
	ID := str.StringToInt(reqID)

	input := new(requests.CakeRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.CakeUC{ContractUC: h.ContractUC}
	err := uc.EditCake(c, ID, input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err.Error(), http.StatusBadRequest)
	}

	return h.SendResponse(ctx, "success", nil, nil, 0)
}
