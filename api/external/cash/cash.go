package cash

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	ReqAddr      = "req:"
	StartGetSign = "GET-CASH-START"
	FinGetSign   = "GET-CASH-[%s]"

	StartSetSign = "SET-CASH-START-[%s]"
	FinSetSign   = "SET-CASH-[%s]"
)

type CashStorage struct {
	Client *redis.Client
}

func NewCashStorage(port string, ctx context.Context) (*CashStorage, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + port,
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &CashStorage{
		Client: rdb,
	}, nil
}
