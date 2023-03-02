package graceful

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	shutdownTimeout = 10*time.Second
)

type ShutdownFunc func() error

var (
	handler   *shutdownHandler
	execOnErr func(error)
)

var (
	ErrTimeoutExceeded = errors.New("graceful shutdown failed: timeout exceeded")
	ErrForceShutdown = errors.New("graceful shutdown failed: force shutdown")
)

func setupHandler() {
	notify := make(chan os.Signal, 1)
	forceStop := make(chan struct{}, 1)
	signal.Notify(notify, syscall.SIGINT, syscall.SIGTERM)
	handler = newHandler(notify, forceStop)

	execOnErr = func(err error) {
		log.Printf("shutdown callback error: %v", err)
	}
}

func init() {
	setupHandler()
}

func AddCallback(fn ShutdownFunc) {
	handler.add(fn)
}

func ShutdownNow() {
	handler.forceStop <- struct{}{}
}

func WaitShutdown() error {
	select {
	case <-handler.stop:
	case <-handler.forceStop:
	}

	handler.markAsShutdown()

	notify := make(chan os.Signal, 1)
	signal.Notify(notify, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := len(handler.callbacks) - 1; i >= 0; i-- {
			err := handler.callbacks[i]()
			if err != nil && execOnErr != nil {
				execOnErr(err)
			}
		}
	}()

	select {
	case <-done:
		return nil
	case <-notify:
		return ErrForceShutdown
	case <-ctx.Done():
		return ErrTimeoutExceeded
	}
}
