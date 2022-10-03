package app

import (
	"context"
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

func (f *FileServerApp) UploadFileToServer(ctx context.Context, files []*multipart.FileHeader, name, description string) error {
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
		Description: description,
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
	errChan := make(chan error)
	done := make(chan struct{})

	go func() {
		err := f.FileServerStorage.DeleteFile(fileName)
		if err != nil {
			f.Logger.Error("Error delete file from server", zap.Error(err))
			errChan <- err
		}

		close(done)
	}()

	go func() {
		<-done
		err := f.Storage.DeleteRecord(context.Background(), id)
		if err != nil {
			f.Logger.Error("Error delete record from database", zap.Error(err))
			errChan <- err
		}

		close(errChan)
	}()

	select {
	case err := <-errChan:
		return err
	}
}
