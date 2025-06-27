package logger

import (
	"log"
	"log/slog"
)

// safeLogError логирует ошибку с защитой от паники
func safeLogError(logger *slog.Logger, msg string, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in slog logging: %v", r)
		}
	}()

	if err != nil {
		logger.Error(msg, "error", err.Error())
	} else {
		logger.Error(msg)
	}
}

// safeLogInfo логирует информационное сообщение с защитой от паники
func safeLogInfo(logger *slog.Logger, msg string, args ...any) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in slog info logging: %v", r)
		}
	}()

	logger.Info(msg, args...)
}
