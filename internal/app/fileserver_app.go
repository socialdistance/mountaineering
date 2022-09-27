package app

import (
	"context"
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

func (f *FileServerApp) UploadFileToServer(ctx context.Context, files []*multipart.FileHeader, name string) {
	//opCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	//defer cancel()

	// fix this
	errorCh := make(chan error)
	resultCh := make(chan string)

	go func() {
		resultCh, errorCh = f.File.CreateFile(files)
	}()
	defer close(resultCh)
	defer close(errorCh)

	select {
	case result := <-resultCh:
		fmt.Println(result)
	case err := <-errorCh:
		fmt.Println(err)
	}
}
