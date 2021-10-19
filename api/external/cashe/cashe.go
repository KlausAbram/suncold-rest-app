package cashe

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

type CasheStorage struct {
	Client *redis.Client
}

func NewCasheStorage(port string, ctx context.Context) (*CasheStorage, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + port,
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &CasheStorage{
		Client: rdb,
	}, nil
}
