package redis

import (
	"context"
	"strconv"

	goredis "github.com/redis/go-redis/v9"
	"github.com/sahar-mirtalebi/quiz-battle/internal/category"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/timestamp"
)

func (db *MatchingDB) AddToWaitingList(ctx context.Context, userID uint, category category.Category) error {

	_, err := db.Adapter.Client().ZAdd(ctx, db.getWaitingListKey(string(category)), goredis.Z{
		Score:  float64(timestamp.Now()),
		Member: strconv.Itoa(int(userID)),
	}).Result()
	if err != nil {
		return richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
