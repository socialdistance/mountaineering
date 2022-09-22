package app

import (
	"go.uber.org/zap"
)

type App struct {
	logger  Logger
	storage Storage
}

type Storage interface {
}

type Logger interface {
	Debug(message string, fields ...zap.Field)
	Info(message string, fields ...zap.Field)
	Error(message string, fields ...zap.Field)
	Fatal(message string, fields ...zap.Field)
	Sync() error
}

func NewApp(logger Logger, storage Storage) *App {
	return &App{logger: logger, storage: storage}
}
