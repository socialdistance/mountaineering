package http

import (
	"github.com/labstack/echo/v4"
	"mountaineering/internal/app"
)

type Router struct {
	app    *app.App
	logger Logger
}

func NewRouters(app *app.App, logger Logger) *Router {
	return &Router{app: app, logger: logger}
}

func (r *Router) CreateServiceRouter(c echo.Context) (err error) {

	return err
}

func (r *Router) DeleteServiceRouter(c echo.Context) (err error) {

	return err
}

func (r *Router) UpdateServiceRouter(c echo.Context) (err error) {

	return err
}
