package redis

import (
	redisadapter "github.com/sahar-mirtalebi/quiz-battle/internal/adapter/redis"
)

type Config struct {
	WaitingListPrefix string `koanf:"waiting_list_prefix"`
}

type MatchingDB struct {
	config  Config
	Adapter redisadapter.Adapter
}

func New(adapter redisadapter.Adapter, config Config) *MatchingDB {
	return &MatchingDB{Adapter: adapter, config: config}
}

func (db *MatchingDB) getWaitingListKey(category string) string {
	return db.config.WaitingListPrefix + ":" + category
}
