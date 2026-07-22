// Package matchingrepo stores matchmaking queues in Redis.
package matchingrepo

import (
	goredis "github.com/redis/go-redis/v9"
	redisadapter "github.com/sahar-mirtalebi/quiz-battle/internal/adapter/redis"
)

type MatchingDB struct {
	db goredis.Cmdable
}

func New(db redisadapter.Adapter) *MatchingDB {
	return &MatchingDB{db: db.Client()}
}
