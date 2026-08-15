package redis

import (
	redisadapter "github.com/sahar-mirtalebi/quiz-battle/internal/adapter/redis"
)

type MatchingDB struct {
	Adapter redisadapter.Adapter
}

func New(adapter redisadapter.Adapter) *MatchingDB {
	return &MatchingDB{Adapter: adapter}
}
