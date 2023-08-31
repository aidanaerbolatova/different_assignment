package logger

import (
	"log"

	"go.uber.org/zap"
)

func Logger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error while create logger", err)
		return nil, err
	}
	defer logger.Sync()
	log := logger.Sugar()

	return log, nil
}
