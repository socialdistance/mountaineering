package app

import (
	"context"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"mime/multipart"
	internalstorage "mountaineering/internal/storage"
	"time"
)

type FileServerApp struct {
	Logger            Logger
	Storage           Storage
	FileServerStorage FileServerStorage
}

type FileServerStorage interface {
	CreateFile([]*multipart.FileHeader, chan string, chan error) chan error
}

func NewFileServerApp(logger Logger, storage Storage) *FileServerApp {
	f := internalstorage.NewFileServerStorage()

	return &FileServerApp{
		Logger:            logger,
		Storage:           storage,
		FileServerStorage: f,
	}
}

func (f *FileServerApp) UploadFileToServer(ctx context.Context, files []*multipart.FileHeader, name string) error {
	opCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	resultCh := make(chan string, 1)
	errCh := make(chan error)
	defer close(resultCh)
	defer close(errCh)

	var res string

	go func() {
		errCh = f.FileServerStorage.CreateFile(files, resultCh, errCh)
	}()

	select {
	case err := <-errCh:
		//f.Logger.Error("error create file in fs", zap.Error(err))
		return err
	case res = <-resultCh:
	}

	file := internalstorage.FileServer{
		ID:          uuid.FromStringOrNil("test"),
		Name:        name,
		Path:        res,
		Description: "test",
	}

	err := f.Storage.CreateRecordForFile(opCtx, file)
	if err != nil {
		f.Logger.Error("error create record in database", zap.Error(err))
		return err
	}

	return nil
}
