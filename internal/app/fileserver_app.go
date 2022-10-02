package app

import (
	"context"
	"fmt"
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
	CreateFile([]*multipart.FileHeader, chan string) error
	DeleteFile(file string) error
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

	resultCh := make(chan string)
	defer close(resultCh)

	errorCh := make(chan error)
	defer close(errorCh)

	var res string

	go func() {
		err := f.FileServerStorage.CreateFile(files, resultCh)
		if err != nil {
			errorCh <- err
		}
	}()

	select {
	case res = <-resultCh:
	case err := <-errorCh:
		f.Logger.Error("Error upload file", zap.Error(err))
		return err
	}

	file := internalstorage.FileServer{
		Name:        name,
		Path:        res,
		Description: "test",
	}

	err := f.Storage.CreateRecordForFile(opCtx, file)
	if err != nil {
		f.Logger.Error("error create record in database", zap.Error(err))
		return err
	}
	f.Logger.Info("[+] Success upload file with name and path", zap.Strings("file", []string{name, res}))

	return nil
}

func (f *FileServerApp) DeleteFileFromServer(ctx context.Context, id, fileName string) error {
	opCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	//doneCh := make(chan struct{})

	// Я хочу чтобы фукнция на удаление из файла и на удаление из базы работали параллельно.

	go func() {
		err := f.FileServerStorage.DeleteFile(fileName)
		fmt.Println("err1", err)
	}()

	go func() {
		err := f.Storage.DeleteRecord(opCtx, id)
		fmt.Println("err2", err)
		// err2 -> context canceled
	}()

	return nil
}
