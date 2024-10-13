package logging

import (
	"go.uber.org/zap"
	"log"
)

var logger *zap.SugaredLogger

func InitLogger() {
	// init logger
	if logger != nil {
		return
	}
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Panicln("failed to init logger", err)
	}
	logger = l.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		log.Println("logger is not initialized")
		InitLogger()
	}
	return logger
}
