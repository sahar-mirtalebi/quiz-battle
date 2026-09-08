// Package redis provides the application's Redis adapter.
package redis

import (
	"context"
	"net"
	"strconv"

	goredis "github.com/redis/go-redis/v9"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

type Config struct {
	Host string `koanf:"host"`
	Port int    `koanf:"port"`
	DB   int    `koanf:"db"`
}

type Adapter struct {
	client *goredis.Client
}

func New(ctx context.Context, cfg Config) Adapter {
	client := goredis.NewClient(&goredis.Options{
		Addr: net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		DB:   cfg.DB,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		panic(richerror.New().WithError(err).WithKind(richerror.KindUnexpected).WithMessage("redis connection failed"))
	}

	return Adapter{
		client: client,
	}
}

func (a Adapter) Client() goredis.Cmdable {
	return a.client
}

func (a Adapter) Close() error {
	if err := a.client.Close(); err != nil {
		return richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
