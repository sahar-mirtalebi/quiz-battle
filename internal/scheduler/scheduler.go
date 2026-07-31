package scheduler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type Scheduler struct {
}

func New() Scheduler {
	return Scheduler{}
}

func (s Scheduler) Start(ctx context.Context) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	fmt.Println("start")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("scheduler stopped")
			return
		case <-ticker.C:
			if err := s.runJob(ctx); err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}

				log.Println("job failed:", err)
			}
		}

	}
}

func (s Scheduler) runJob(ctx context.Context) error {
	return nil
}
