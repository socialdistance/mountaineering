package fileserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	internalapp "mountaineering/internal/app"
	"net/http"
	"os"
)

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

func (r *RouterFileServer) Upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(fmt.Sprintf("./uploads/%s", file.Filename))
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and email=%s.</p>", len(files), name, email))
}
