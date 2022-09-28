package app

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
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
	CreateFile([]*multipart.FileHeader, chan string) chan error
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

	// fix this
	//errorCh := make(chan error)

	//go func() {
	//defer close(errorCh)

	doneCh := make(chan struct{})
	resultCh := make(chan string)
	var file internalstorage.FileServer

	go func() {
		err := f.FileServerStorage.CreateFile(files, resultCh)
		if err != nil {
			fmt.Println("err", err)
		}

		close(doneCh)
	}()

	go func() {
		<-doneCh
		res := <-resultCh

		file.ID = uuid.FromStringOrNil("test")
		file.Name = res
		file.Path = res
		file.Description = "test"

		err := f.Storage.CreateRecordForFile(opCtx, file)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(err)
	}()
	//defer close(doneCh)
	//defer close(resultCh)

	return nil
}
