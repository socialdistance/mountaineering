package http

import (
	"mountaineering/internal/app"
)

type Router struct {
	app    *app.App
	logger Logger
}

func NewRouters(app *app.App, logger Logger) *Router {
	return &Router{app: app, logger: logger}
}
