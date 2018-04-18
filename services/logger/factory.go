package logger

import (
	"go.uber.org/zap"
)

func GetSugaredLogger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	sl := logger.Sugar()
	return sl
}
