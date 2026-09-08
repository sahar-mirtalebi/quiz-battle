package redis

import (
	"context"
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func (d DB) Upsert(ctx context.Context, key string, timestamp uint64, expirationTime time.Duration) error {
	_, err := d.Adapter.Client().Set(ctx, key, timestamp, expirationTime).Result()
	if err != nil {
		return richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
