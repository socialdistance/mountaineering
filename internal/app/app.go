package app

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	internalstorage "mountaineering/internal/storage"
	"time"
)

type App struct {
	logger  Logger
	storage Storage
}

type Storage interface {
	CreateRecordForFile(ctx context.Context, file internalstorage.FileServer) error
	DeleteRecord(ctx context.Context, id string) error

	CreateService(ctx context.Context, service internalstorage.Services) error
	DeleteService(ctx context.Context, id string) error
	UpdateService(ctx context.Context, m map[string]interface{}) error
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

func (a *App) CreateServiceApp(ctx context.Context, service internalstorage.Services) error {
	opCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := a.storage.CreateService(opCtx, service)
	if err != nil {
		a.logger.Error("Can't create service", zap.Error(err))
		return err
	}

	return err
}

func (a *App) DeleteServiceApp(ctx context.Context, id string) error {
	opCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := a.storage.DeleteService(opCtx, id)
	if err != nil {
		a.logger.Error("Can't delete service", zap.Error(err))
		return err
	}

	return err
}

func (a *App) UpdateServiceApp(ctx context.Context, service internalstorage.Services) error {
	opCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var m map[string]interface{}

	marshal, err := json.Marshal(service)
	if err != nil {
		a.logger.Error("Cant marshal service", zap.Error(err))
		return err
	}

	err = json.Unmarshal(marshal, &m)
	if err != nil {
		a.logger.Error("Cant unmarshall service", zap.Error(err))
		return err
	}

	err = a.storage.UpdateService(opCtx, m)
	if err != nil {
		a.logger.Error("Can't update service", zap.Error(err))
		return err
	}

	return err
}
