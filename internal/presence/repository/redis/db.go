package redis

import (
	redis "github.com/sahar-mirtalebi/quiz-battle/internal/adapter/redis"
)

type DB struct {
	Adapter redis.Adapter
}

func New(adapter redis.Adapter) *DB {
	return &DB{Adapter: adapter}
}
