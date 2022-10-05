package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mountaineering/internal/app"
	"mountaineering/internal/server"
	"net/http"
)

type Router struct {
	app    *app.App
	logger Logger
}

func NewRouters(app *app.App, logger Logger) *Router {
	return &Router{app: app, logger: logger}
}

func (r *Router) CreateServiceRouter(c echo.Context) (err error) {
	dto := new(server.ServiceDto)

	if err = c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, server.HTTPError{Error: "Cant bind data for create service"})
	}

	dtoModel := dto.GetModelCreateService()

	err = r.app.CreateServiceApp(c.Request().Context(), *dtoModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, server.HTTPError{Error: "Cant create service"})
	}

	return c.JSON(http.StatusOK, server.HTTPSuccess{Success: "success"})
}

func (r *Router) DeleteServiceRouter(c echo.Context) (err error) {
	deleteDto := new(server.DeleteServiceDto)

	if err = c.Bind(deleteDto); err != nil {
		return c.JSON(http.StatusBadRequest, server.HTTPError{Error: "Cant bind data for delete service"})
	}

	err = r.app.DeleteServiceApp(c.Request().Context(), deleteDto.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, server.HTTPError{Error: "Cant delete service"})
	}

	return c.JSON(http.StatusOK, server.HTTPSuccess{Success: "success"})
}

func (r *Router) UpdateServiceRouter(c echo.Context) (err error) {
	dto := new(server.ServiceDto)

	if err = c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, server.HTTPError{Error: "Cant bind data for update service"})
	}

	dtoModel := dto.GetModelCreateService()
	fmt.Printf("%+v", dtoModel)

	err = r.app.UpdateServiceApp(c.Request().Context(), *dtoModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, server.HTTPError{Error: "Cant update service"})
	}

	return c.JSON(http.StatusOK, server.HTTPSuccess{Success: "success"})
}
