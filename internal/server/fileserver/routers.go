package fileserver

import internalapp "mountaineering/internal/app"

type RouterFileServer struct {
	logger internalapp.Logger
	app    internalapp.FileServerApp
}

func NewRouterFileServer(fileServerApp internalapp.FileServerApp, logger internalapp.Logger) *RouterFileServer {
	return &RouterFileServer{
		logger: logger,
		app:    fileServerApp,
	}
}
