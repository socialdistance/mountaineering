package app

import (
	"context"
	"go.uber.org/zap"
	"mime/multipart"
	"mountaineering/internal/storage"
)

type App struct {
	logger  Logger
	storage Storage
}

type Storage interface {
	CreateRecordForFile(ctx context.Context, file storage.File) error
}

type File interface {
	CreateFile(files []*multipart.FileHeader) (chan string, chan error)
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
