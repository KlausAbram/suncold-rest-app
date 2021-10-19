package usecase

import (
	"context"

	"github.com/klaus-abram/suncold-restful-app/api/external/cashe"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type CasheCase struct {
	rdb *cashe.CasheStorage
}

func NewCasheStorage(rdb *cashe.CasheStorage) *CasheCase {
	return &CasheCase{
		rdb: rdb,
	}
}

func (chc *CasheCase) GetCashedRequests(ctx context.Context) (*[]models.WeatherRequest, error) {
	data, err := chc.rdb.GetAllCashedRequests(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
