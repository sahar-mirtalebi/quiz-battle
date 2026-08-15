// Command scheduler starts the background jobs process.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	redisadapter "github.com/sahar-mirtalebi/quiz-battle/internal/adapter/redis"
	"github.com/sahar-mirtalebi/quiz-battle/internal/config"
	matchingrepo "github.com/sahar-mirtalebi/quiz-battle/internal/matching/repository/redis"
	matchingservice "github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"
	"github.com/sahar-mirtalebi/quiz-battle/internal/scheduler"
)

func main() {
	cfg := config.Load("config.yml")

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	startupCtx, cancelStartup := context.WithTimeout(context.Background(), 10*time.Second)
	cache := redisadapter.New(startupCtx, cfg.Redis)
	cancelStartup()
	defer cache.Close()

	matchingRepo := matchingrepo.New(cache)
	matchingService := matchingservice.New(matchingRepo, cfg.MatchingService)

	cronScheduler, err := scheduler.New(
		matchingService,
		scheduler.Config{
			MatchPlayersInterval: cfg.Scheduler.MatchPlayersInterval,
		},
	)
	if err != nil {
		log.Fatalf("create scheduler: %v", err)
	}

	cronScheduler.Start()
	log.Println("scheduler started")

	<-ctx.Done()
	log.Println("shutting down scheduler")

	shutdownCtx, cancelShutdown := context.WithTimeout(
		context.Background(),
		cfg.Application.GracefullShutdownTimeout,
	)
	defer cancelShutdown()

	if err := cronScheduler.Shutdown(shutdownCtx); err != nil {
		log.Printf("scheduler shutdown failed: %v", err)
	}
}
