package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type MatchingService interface {
	MatchPlayers(ctx context.Context) error
}

type Scheduler struct {
	scheduler       gocron.Scheduler
	matchingService MatchingService
}

func New(matchingService MatchingService, interval time.Duration) (*Scheduler, error) {
	cronScheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	s := &Scheduler{
		scheduler:       cronScheduler,
		matchingService: matchingService,
	}

	if err := s.registerJobs(interval); err != nil {
		_ = cronScheduler.Shutdown()
		return nil, err
	}

	return s, nil
}

func (s *Scheduler) registerJobs(interval time.Duration) error {
	_, err := s.scheduler.NewJob(
		gocron.DurationJob(interval),
		gocron.NewTask(func(ctx context.Context) {
			if err := s.matchingService.MatchPlayers(ctx); err != nil {
				log.Printf("match players job failed: %v", err)
			}
		}),
		gocron.WithName("match-players"),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)

	return err
}

func (s *Scheduler) Start() {
	s.scheduler.Start()
}

func (s *Scheduler) Shutdown(ctx context.Context) error {
	return s.scheduler.ShutdownWithContext(ctx)
}
