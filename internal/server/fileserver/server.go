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

	//e.Use(middleware.Static("/home/user/work/mountaineering/uploads"))

	return &FileServer{
		host:   host,
		port:   port,
		e:      e,
		router: router,
	}
}

// BuildRouters TODO: serve static files
func (f *FileServer) BuildRouters() {
	f.e.Static("/", "uploads")
	fs := http.FileServer(http.Dir("/home/user/work/mountaineering/uploads"))
	f.e.GET("/uploads/*", echo.WrapHandler(http.StripPrefix("/uploads/", fs)))

	fsAPI := f.e.Group("/api")

	fsAPI.POST("/upload", f.router.Upload)
	fsAPI.DELETE("/delete", f.router.Delete)
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
