package app

type FileServerApp struct {
	Logger            Logger
	StorageFileServer StorageFileServer
}

type StorageFileServer interface {
}

func NewFileServerApp(logger Logger, storage StorageFileServer) *FileServerApp {
	return &FileServerApp{
		Logger:            logger,
		StorageFileServer: storage,
	}
}
