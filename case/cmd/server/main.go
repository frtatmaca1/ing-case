package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/frtatmaca/case/config"
	"github.com/frtatmaca/case/controller"
	"github.com/frtatmaca/case/middleware"
	"github.com/frtatmaca/case/pkg/logging"
	"github.com/frtatmaca/case/pkg/server"
	proxy "github.com/frtatmaca/case/proxy/person"
	"github.com/frtatmaca/case/service"
)

func main() {
	cfg := config.NewConfiguration()

	logger := logging.NewLoggerWithLevel("stderr", cfg.LogLevel)
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Fatalf("error syncing logs:  %v", err)
		}
	}()

	httpClient := http.DefaultClient
	httpClient.Timeout = cfg.PersonApi.Timeout

	personClient := proxy.NewClient(cfg.PersonApi.BaseUrl, httpClient)
	personService := service.NewPersonService(personClient)
	personController := controller.NewPersonController(personService)

	srv := server.NewHTTPServer(cfg.Port)

	srv.Engine.Use(
		middleware.TracingHandler(),
		middleware.ErrorHandler(logger),
	)

	groupV1 := srv.Engine.Group("/api/v1")
	{
		contentGroup := groupV1.Group("/persons")
		{
			contentGroup.GET("", personController.Person)
		}
	}

	chanErrors := make(chan error)
	go func() {
		chanErrors <- srv.Run()
	}()

	chanSignals := make(chan os.Signal, 1)
	signal.Notify(chanSignals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-chanErrors:
		logger.Fatal("Unable to run server", err.Error())
	case s := <-chanSignals:
		logger.Infof("Warning: Received %s signal, aborting in 5 seconds...", s)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Close(ctx); err != nil {
			logger.Fatal("Server forced to shutdown: ", err)
		}
		logger.Info("Server exiting gracefully")
	}
}
