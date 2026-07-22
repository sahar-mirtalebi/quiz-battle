package matchingrepo

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
)

func (db *MatchingDB) AddToWaitingList(ctx context.Context, userID uint, category quizdomain.Category, joinedAt time.Time) error {
	key := fmt.Sprintf("waiting-list:%s", category)

	return db.db.ZAdd(ctx, key, goredis.Z{
		Score:  float64(joinedAt.UnixMilli()),
		Member: userID,
	}).Err()
}
