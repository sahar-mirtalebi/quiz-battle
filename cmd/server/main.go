// Command server starts the quiz game API.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/mysql/migrator"
	"github.com/sahar-mirtalebi/quiz-battle/internal/app"
	"github.com/sahar-mirtalebi/quiz-battle/internal/config"
)

func main() {
	cfg := config.Load("config.yml")
	startupCtx, cancelStartup := context.WithTimeout(context.Background(), 10*time.Second)

	// TODO - add command for migration
	migrator := migrator.New(cfg.Mysql)
	migrator.Up(startupCtx)

	application, err := app.New(startupCtx, *cfg)
	cancelStartup()
	if err != nil {
		log.Fatalf("create application: %v", err)
	}
	defer func() {
		if err := application.Close(); err != nil {
			log.Printf("close application resources: %v", err)
		}
	}()

	application.Server.Serve()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	log.Println("received shutdown signal")

	shutDownCtx, cancel := context.WithTimeout(
		context.Background(),
		cfg.Application.GracefullShutdownTimeout,
	)
	defer cancel()

	if err := application.Server.Shutdown(shutDownCtx); err != nil {
		log.Printf("http server shutdown error: %v", err)
	}
}
