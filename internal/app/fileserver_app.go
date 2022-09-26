package app

import (
	"context"
	"mime/multipart"
	"mountaineering/internal/storage"
	"time"
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
	opCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var data storage.File

	doneChannel := make(chan struct{})
	errorChannel := make(chan error)

	go func(done chan struct{}, error chan error) {
		path, err := f.File.CreateFile(files)
		if err != nil {
			error <- err
		}

		done <- struct{}{}
	}(doneChannel, errorChannel)

	//go func() {
	//	f.StorageFileServer.CreateRecordForFile(opCtx)
	//}()

}
