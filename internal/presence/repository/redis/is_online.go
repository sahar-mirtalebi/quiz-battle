package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func (d DB) IsOnline(ctx context.Context, keys []string) (map[string]uint64, error) {
	existing := make(map[string]uint64, len(keys))
	if len(keys) == 0 {
		return existing, nil
	}

	values, err := d.Adapter.Client().MGet(ctx, keys...).Result()
	if err != nil {
		return nil, richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}
	for i, key := range keys {
		if values[i] == nil {
			continue
		}
		value, err := strconv.ParseUint(fmt.Sprint(values[i]), 10, 64)
		if err != nil {
			return nil, richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
		}
		existing[key] = value
	}
	return existing, nil
}
