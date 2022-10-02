package fileserver

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type FileServer struct {
	host   string
	port   string
	e      *echo.Echo
	router *RouterFileServer
}

func NewFileServer(host, port string, router *RouterFileServer) *FileServer {
	e := echo.New()
	e.HideBanner = true

	return &FileServer{
		host:   host,
		port:   port,
		e:      e,
		router: router,
	}
}

func (f *FileServer) BuildRouters() {
	f.e.Static("/", "uploads")

	f.e.POST("/upload", f.router.Upload)
	f.e.DELETE("/delete", f.router.Delete)
}

func (f *FileServer) Start() error {
	if err := f.e.Start(fmt.Sprintf(":%s", f.port)); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server stopped: %w", err)
	}

	return nil
}

func (f *FileServer) Stop(ctx context.Context) error {
	optCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	if err := f.e.Shutdown(optCtx); err != nil {
		return fmt.Errorf("could not shutdown server gracefuly: %w", err)
	}

	return nil
}
