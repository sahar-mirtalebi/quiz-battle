// Command server starts the quiz game API.
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sahar-mirtalebi/quiz-battle/internal/scheduler"
)

func main() {

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	scheduler := scheduler.New()
	scheduler.Start(ctx)

}
