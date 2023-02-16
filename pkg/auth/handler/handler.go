package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/flash-card-api/db"
	"github.com/minhquy1903/flash-card-api/server"
)

type Handler struct {
	c *context.Context
	db server.Server
}

func Register(c echo.Context, svr server.Server) error {
	
	return nil
}
