package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/dto"
)

type MatchingService interface {
	MatchPlayers(ctx context.Context, req dto.MatchPlayersInput) (dto.MatchPlayersOutput, error)
}

type Config struct {
	MatchPlayersInterval time.Duration
}

type Scheduler struct {
	scheduler       gocron.Scheduler
	matchingService MatchingService
}

func New(matchingService MatchingService, cfg Config) (*Scheduler, error) {
	cronScheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	s := &Scheduler{
		scheduler:       cronScheduler,
		matchingService: matchingService,
	}

	if err := s.registerJobs(cfg); err != nil {
		_ = cronScheduler.Shutdown()
		return nil, err
	}

	return s, nil
}

func (s *Scheduler) registerJobs(cfg Config) error {
	if err := s.registerMatchPlayersJob(cfg.MatchPlayersInterval); err != nil {
		return err
	}

	return nil
}

func (s *Scheduler) registerMatchPlayersJob(interval time.Duration) error {
	_, err := s.scheduler.NewJob(
		gocron.DurationJob(interval),
		gocron.NewTask(func(ctx context.Context) {
			if _, err := s.matchingService.MatchPlayers(ctx, dto.MatchPlayersInput{}); err != nil {
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
