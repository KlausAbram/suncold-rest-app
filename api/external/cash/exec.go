package cash

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/klaus-abram/suncold-restful-app/models"
	"github.com/sirupsen/logrus"
)

var (
	ErrParseResCash  = "eror with parsing result by get - [%s]"
	ErrSetCash       = "error with setting cash key-value - [%s]"
	ErrGetCash       = "error with getting cash key-value - [%s]"
	ErrUnmarshalCash = "eror with parsing json to model - [%s]"
)

func (csh *CashStorage) SetRequestToCash(req *models.WeatherRequest, ctx context.Context) error {
	jsonVal, err := json.Marshal(*req)
	if err != nil {
		return err
	}

	id := ReqAddr + strconv.Itoa(req.Id)
	logrus.Printf(StartSetSign, id)

	if err := csh.Client.Set(ctx, id, jsonVal, 0).Err(); err != nil {
		logrus.Errorf(ErrSetCash, err.Error())
		return err
	}

	logrus.Printf(FinSetSign, id)

	return nil
}

func (csh *CashStorage) GetAllCashedRequests(ctx context.Context) (*[]models.WeatherRequest, error) {
	var input []models.WeatherRequest

	logrus.Printf(StartGetSign)

	iter := csh.Client.Scan(ctx, 0, "", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		val, err := csh.Client.Get(ctx, key).Result()
		if err != nil {
			logrus.Errorf(ErrParseResCash, err.Error())
			return nil, err
		}

		data := []byte(val)
		var req = models.WeatherRequest{}

		if err := json.Unmarshal(data, &req); err != nil {
			logrus.Errorf(ErrUnmarshalCash, err.Error())
			return nil, err
		}
		logrus.Printf(fmt.Sprintf(FinGetSign, key))

		input = append(input, req)
	}

	if err := iter.Err(); err != nil {
		logrus.Fatalf("%s - [%s]", "error with iteration", err.Error())
		return nil, err
	}

	return &input, nil
}
