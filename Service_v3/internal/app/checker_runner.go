package app

import (
	"context"
	"service/internal/app/processors"
	"service/internal/checker"
	"service/internal/checker/gometrclient"
	"service/internal/checker/googlemetr"
)

func runChecker(server *Server, processor *processors.MetricsProcessor) error {
	// Create a new checker
	che := new(checker.Checker)

	// Create a channel with a buffer to send checkables
	ch := make(chan checker.Checkable, 2)

	// Create cancel context
	_, cancel := context.WithCancel(context.Background())

	// Add Clients with the channel to the checker
	ch <- gometrclient.NewClient(processor)
	go che.Add(ch)

	ch <- googlemetr.NewClient(processor)
	go che.Add(ch)

	// Schedule checker to run every 30 seconds
	server.scheduler.Every(30).Seconds().Do(func() {
		go che.Run()
	})

	// Cancel the checker when the context is done
	che.Stop(cancel)

	return nil
}
