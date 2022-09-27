package app

import (
	"context"
	"fmt"
	"mime/multipart"
)

type FileServerApp struct {
	Logger            Logger
	StorageFileServer Storage
}

type FileServer interface {
	CreateFile(files []*multipart.FileHeader) error
}

func NewFileServerApp(logger Logger, storage Storage) *FileServerApp {
	return &FileServerApp{
		Logger:            logger,
		StorageFileServer: storage,
	}
}

func (f *FileServerApp) UploadFileToServer(ctx context.Context, files []*multipart.FileHeader, name string) error {
	//opCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	//defer cancel()

	// fix this
	//errorCh := make(chan error)

	//go func() {
	//defer close(errorCh)
	var b FileServer

	err := b.CreateFile(files)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
