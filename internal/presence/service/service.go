package presence

import (
	"context"
	"time"
)

type Config struct {
	ExpirationTime time.Duration `koanf:"expiration_time"`
	PresencePrefix string        `koanf:"prefix"`
}

type Repo interface {
	Upsert(ctx context.Context, key string, timestamp uint64, expirationTime time.Duration) error
}

type Service struct {
	Config Config
	Repo   Repo
}

func New(Config Config, repo Repo) Service {
	return Service{
		Repo:   repo,
		Config: Config,
	}
}
