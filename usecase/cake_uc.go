package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/mini-project/db/repository"
	"github.com/mini-project/db/repository/models"
	"github.com/mini-project/pkg/functioncaller"
	"github.com/mini-project/pkg/logger"
	"github.com/mini-project/usecase/requests"
)

type CakeUC struct {
	*ContractUC
	*sql.Tx
}

//BuildBody ...
func (uc CakeUC) BuildBody(res *models.Cake) {}

func (uc CakeUC) AddCake(c context.Context, input *requests.CakeRequest) (res models.Cake, err error) {
	timeNow := time.Now()
	repo := repository.NewCakeRepository(uc.DB, uc.TX)
	res = models.Cake{
		Title:       input.Title,
		Description: input.Description,
		Rating:      input.Rating,
		Image:       input.Image,
		CreatedAt:   &timeNow,
		UpdatedAt:   &timeNow,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logger.Log(logger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query")
		return res, err
	}

	return res, err
}

func (uc CakeUC) FindByID(c context.Context, ID string) (res models.Cake, err error) {
	repo := repository.NewCakeRepository(uc.DB, uc.TX)
	res, err = repo.FindByID(c, ID)
	if err != nil {
		logger.Log(logger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-find-id")
		return res, err
	}

	uc.BuildBody(&res)

	return res, err
}

func (uc CakeUC) FindAllCake(c context.Context, param models.CakeParameter) (res []models.Cake, err error) {
	repo := repository.NewCakeRepository(uc.DB, uc.TX)
	res, err = repo.FindAllCake(c, param)
	if err != nil {
		logger.Log(logger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-find-all")
		return res, err
	}

	for i := range res {
		uc.BuildBody(&res[i])
	}

	return res, err
}
