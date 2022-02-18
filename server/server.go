package server

import (
	"context"
	"github.com/angelorc/cosmos-balance/client"
	_ "github.com/angelorc/cosmos-balance/swagger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"time"
)

type Server struct {
	*echo.Echo
	client *client.Client
	logger *zap.Logger
}

// @title Cosmos Tracker Server API
// @version 1.0
// @description The cosmos tracker rest server.

func NewServer(client *client.Client, logger *zap.Logger) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	s := &Server{
		Echo:   e,
		client: client,
		logger: logger,
	}
	s.registerRoutes()

	return s
}

func (s *Server) ShutdownWithTimeout(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return s.Shutdown(ctx)
}
