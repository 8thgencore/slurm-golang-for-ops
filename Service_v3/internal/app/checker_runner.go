package app

import (
	"context"
	"service/internal/app/processors"
	"service/internal/checker"
	"service/internal/checker/gometrclient"
)

func runChecker(ctx context.Context, server *Server, processor *processors.MetricsProcessor) error {
	// Create a new checker
	che := new(checker.Checker)

	// Create a channel with a buffer to send checkables
	ch := make(chan checker.Checkable, 1)

	// Add GoMetrClient with the channel to the checker
	goMetrClient := gometrclient.NewGoMetrClient(processor)
	ch <- goMetrClient
	go che.Add(ch)

	// Schedule checker to run every 30 seconds
	server.scheduler.Every(30).Seconds().Do(func() {
		go che.Run()
	})

	// Cancel the checker when the context is done
	_, cancel := context.WithCancel(context.Background())
	che.Stop(cancel)

	return nil
}
