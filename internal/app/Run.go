package app

import "net/http"

func (a *App) Run() error {

	// go a.startWorkers()

	a.Server = &http.Server{
		Addr:    ":" + a.Config.Port,
		Handler: a.Router,
	}

	go func() {

		if err := a.Server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			a.Logger.Error(
				"server error",
				"error",
				err,
			)

		}

	}()

	return a.WaitForShutdown()
}
