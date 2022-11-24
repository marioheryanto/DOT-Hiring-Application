package repository

import (
	"context"

	"github.com/marioheryanto/DOT-Hiring-Application/config"
)

func DeleteRedis(keys ...string) {
	redisCli := config.ConnectRedis()
	redisCli.Del(context.Background(), keys...)
}
