package usecase

import (
	"context"

	"github.com/klaus-abram/suncold-restful-app/api/external/cash"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type CashCase struct {
	rdb *cash.CashStorage
}

func NewCashStorage(rdb *cash.CashStorage) *CashCase {
	return &CashCase{
		rdb: rdb,
	}
}

func (chc *CashCase) GetCashedRequests(ctx context.Context) (*[]models.WeatherRequest, error) {
	data, err := chc.rdb.GetAllCashedRequests(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
