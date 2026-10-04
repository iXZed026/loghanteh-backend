package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (a *App) shutdown() error {

	a.Logger.Info("shutting down application...")

	// 1. Stop Workers
	a.Cancel()

	// 2. Shutdown HTTP Server
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		return err
	}

	// 3. Close Mongo
	// if err := a.DB.Client().Disconnect(ctx); err != nil {
	// 	return err
	// }

	// 4. Close Redis
	// if err := a.Redis.Close(); err != nil {
	// 	return err
	// }

	a.Logger.Info("application stopped successfully")

	return nil
}

func (a *App) WaitForShutdown() error {

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	a.Logger.Info("shutdown signal received")

	return a.shutdown()
}
