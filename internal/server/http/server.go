package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	host   string
	port   string
	e      *echo.Echo
	router *Router
}

type Logger interface {
	Debug(message string, fields ...zap.Field)
	Info(message string, fields ...zap.Field)
	Error(message string, fields ...zap.Field)
	Fatal(message string, fields ...zap.Field)
	LogHTTP(r *http.Request, code, length int)
}

func NewServer(host, port string, router *Router) *Server {
	e := echo.New()
	e.HideBanner = true

	return &Server{host: host, port: port, router: router, e: e}
}

func (s *Server) BuildRouters() {
	serviceAPI := s.e.Group("/services")

	serviceAPI.POST("/create", s.router.CreateServiceRouter)
	serviceAPI.DELETE("/delete", s.router.DeleteServiceRouter)
	serviceAPI.PUT("/update", s.router.UpdateServiceRouter)
}

func (s *Server) Start() error {
	if err := s.e.Start(fmt.Sprintf(":%s", s.port)); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server stopped: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	opCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	if err := s.e.Shutdown(opCtx); err != nil {
		return fmt.Errorf("could not shutdown server gracefuly: %w", err)
	}

	return nil
}
