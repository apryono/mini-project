package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/mini-project/db/repository"
	"github.com/mini-project/db/repository/models"
	"github.com/mini-project/helper"
	"github.com/mini-project/pkg/functioncaller"
	"github.com/mini-project/pkg/logger"
	"github.com/mini-project/usecase/requests"
)

type CakeUC struct {
	*ContractUC
	*sql.Tx
}

//BuildBody ...
func (uc CakeUC) BuildBody(res *models.Cake) {
	res.CreatedAt, _ = time.Parse(time.RFC3339, res.CreatedAt.Add(time.Hour*7).Format(time.RFC3339))
	res.UpdatedAt, _ = time.Parse(time.RFC3339, res.UpdatedAt.Add(time.Hour*7).Format(time.RFC3339))
}

func (uc CakeUC) AddCake(c context.Context, input *requests.CakeRequest) (res models.Cake, err error) {
	timeNow := time.Now()
	repo := repository.NewCakeRepository(uc.DB, uc.TX)
	res = models.Cake{
		Title:       input.Title,
		Description: input.Description,
		Rating:      input.Rating,
		Image:       input.Image,
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logger.Log(logger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query")
		return res, err
	}

	return res, err
}

func (uc CakeUC) FindByID(c context.Context, ID int) (res models.Cake, err error) {
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

func (uc CakeUC) EditCake(c context.Context, ID int, input *requests.CakeRequest) (err error) {
	repo := repository.NewCakeRepository(uc.DB, uc.TX)
	res, err := repo.FindByID(c, ID)
	if err != nil && err != sql.ErrNoRows {
		logger.Log(logger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "checking-exist")
		return err
	}

	req := models.Cake{
		ID:          res.ID,
		Title:       input.Title,
		Description: input.Description,
		Rating:      input.Rating,
		Image:       input.Image,
	}

	if res.ID > 0 {
		res.ID, err = repo.Edit(c, req)
		if err != nil && err != sql.ErrNoRows {
			logger.Log(logger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-edit")
			return err
		}
	} else {
		logger.Log(logger.WarnLevel, helper.DataNotFound, functioncaller.PrintFuncName(), "query-edit-no-data")
		return errors.New(helper.DataNotFound)

	}

	return err
}

func (uc CakeUC) DeleteByID(c context.Context, ID int) (err error) {
	repo := repository.NewCakeRepository(uc.DB, uc.TX)
	res, _ := repo.FindByID(c, ID)
	if res.ID > 0 {
		err = repo.Delete(c, res.ID)
		if err != nil {
			logger.Log(logger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-delete")
			return err
		}
	} else {
		logger.Log(logger.WarnLevel, helper.DataNotFound, functioncaller.PrintFuncName(), "query-delete-no-data")
		return errors.New(helper.DataNotFound)
	}

	return err
}
