package fileserver

import (
	"github.com/labstack/echo/v4"
	internalapp "mountaineering/internal/app"
	"mountaineering/internal/server"
	"net/http"
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

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, server.HTTPError{Error: "Cant get data for upload to server"})
	}
	files := form.File["files"]
	err = r.app.UploadFileToServer(c.Request().Context(), files, name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, server.HTTPError{Error: "Cant upload data to server"})
	}

	return c.JSON(http.StatusOK, server.HTTPSuccess{Success: "Success"})
}

func (r *RouterFileServer) Delete(c echo.Context) (err error) {
	deleteDto := new(server.DeleteDto)

	if err = c.Bind(deleteDto); err != nil {
		return c.JSON(http.StatusBadRequest, server.HTTPError{Error: "Cant bind data"})
	}

	err = r.app.DeleteFileFromServer(c.Request().Context(), deleteDto.Id, deleteDto.FileName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, server.HTTPError{Error: "Cant delete record"})
	}

	return c.JSON(http.StatusOK, server.HTTPSuccess{Success: "success"})
}
