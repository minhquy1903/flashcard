package transport

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/flash-card-api/server"
)

func AuthRoute(r *echo.Group, s *server.Server) {
	r.POST("/register", func(ctx echo.Context) error {
		fmt.Println("hello world")
		return ctx.JSON(http.StatusOK, "Hello word")
	})
	r.POST("/login", func(c echo.Context) error { return nil })
}
