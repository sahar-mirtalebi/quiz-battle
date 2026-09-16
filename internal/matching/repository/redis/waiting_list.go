package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/domain"
	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/timestamp"
)

func (db *MatchingDB) AddToWaitingList(ctx context.Context, userID uint, category quizdomain.Category) error {

	_, err := db.Adapter.Client().ZAdd(ctx, db.getWaitingListKey(string(category)), goredis.Z{
		Score:  float64(timestamp.Now()),
		Member: strconv.Itoa(int(userID)),
	}).Result()
	if err != nil {
		return richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}

func (db *MatchingDB) GetWaitingListByCategory(ctx context.Context, category quizdomain.Category) ([]domain.WaitingMember, error) {

	min := fmt.Sprintf("%d", timestamp.Add(-2*time.Hour))
	max := fmt.Sprintf("%d", timestamp.Now())

	list, err := db.Adapter.Client().ZRangeByScoreWithScores(ctx, db.getWaitingListKey(string(category)), &goredis.ZRangeBy{
		Min: min,
		Max: max,
	}).Result()
	if err != nil {
		return nil, richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}

	result := make([]domain.WaitingMember, 0)
	for _, l := range list {
		userID, _ := strconv.Atoi(l.Member.(string))
		result = append(result, domain.WaitingMember{
			UserID: uint(userID),
			JoinAt: timestamp.ToTime(int64(l.Score)),
		})
	}

	return result, nil

}

func (db *MatchingDB) getWaitingListKey(category string) string {
	return db.config.WaitingListPrefix + ":" + category
}
