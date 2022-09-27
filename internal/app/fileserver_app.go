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
	//errorCh := make(chan error)

	//go func() {
	//defer close(errorCh)

	errorCh := f.File.CreateFile(files)
	//}()
	fmt.Println(errorCh)

}
