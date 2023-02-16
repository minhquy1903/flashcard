package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/flash-card-api/config"
	"github.com/minhquy1903/flash-card-api/db"
	"github.com/minhquy1903/flash-card-api/pkg/auth/transport"
	"github.com/minhquy1903/flash-card-api/server"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := db.GetPostgresInstance(cfg)

	svr := server.NewServer(cfg, db, echo.New())

	svr.Echo.GET("/heath", func(c echo.Context) error { return c.JSON(http.StatusOK, "alive") })

	transport.AuthRoute(svr.Echo.Group("/auth"), svr)

	svr.Run()
}
