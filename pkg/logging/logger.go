package logging

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func InitLogger() {
	// init logger
	if logger != nil {
		return
	}
	l, _ := zap.NewDevelopment()
	logger = l.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		panic("logger is not initialized")
	}
	return logger
}
