package app

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
)

type FileServerApp struct {
	Logger            Logger
	StorageFileServer Storage
	File              File
}

func NewFileServerApp(logger Logger, storage Storage) *FileServerApp {
	return &FileServerApp{
		Logger:            logger,
		StorageFileServer: storage,
	}
}

func (f *FileServerApp) UploadFileToServer(ctx context.Context, files []*multipart.FileHeader) {
	//opCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	//defer cancel()

	errorCh := make(chan error, 0)
	resultsCh := make(chan interface{}, 0)

	go func() {
		result, err := f.File.CreateFile(files)
		if err != nil {
			errorCh <- errors.New("Does not compute")
		}

		resultsCh <- result
	}()

	select {
	case err := <-errorCh:
		fmt.Println(err)
	case res := <-resultsCh:
		fmt.Println(res)
	}

}
