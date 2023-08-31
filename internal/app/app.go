package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"rest/config"
	"rest/internal/app/logger"
	"rest/internal/repository"
	"rest/internal/service"
	"rest/internal/transport/httpserver"
	"rest/internal/transport/httpserver/handlers"
)

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gracefullyShutdown(cancel)

	logger, err := logger.Logger()
	if err != nil {
		logger.Errorf("Error while initialization logger: %v", err)
		return err
	}

	config, err := config.ParseConfig()
	if err != nil {
		logger.Errorf("Error while parse config file: %v", err)
		return err
	}
	db, err := repository.ConnectSQL(logger, config)
	if err != nil {
		logger.Errorf("Error while connect to SQL: %v", err)
		return err
	}
	redis, err := repository.ConnectRedis(logger, config)
	if err != nil {
		logger.Errorf("Error while connect to redis: %v", err)
		return err
	}

	repository := repository.NewRepository(db, redis, logger, ctx)
	service := service.NewService(repository, logger)
	handlers := handlers.NewHandler(service)

	srv := new(httpserver.Server)
	if err := srv.Run("8080", handlers.InitRoute()); err != nil {
		log.Printf(err.Error())
	}
	return nil
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
