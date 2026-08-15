package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	goredis "github.com/redis/go-redis/v9"
	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func (db *MatchingDB) AddToWaitingList(ctx context.Context, userID uint, category quizdomain.Category) error {
	key := fmt.Sprintf("waiting-list:%s", category)

	_, err := db.Adapter.Client().ZAdd(ctx, key, goredis.Z{
		Score:  float64(time.Now().UnixMilli()),
		Member: strconv.Itoa(int(userID)),
	}).Result()
	if err != nil {
		return richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
