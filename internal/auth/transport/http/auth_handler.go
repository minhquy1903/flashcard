package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/flashcard-api/internal/auth/presenter"
	"github.com/minhquy1903/flashcard-api/internal/auth/service"
)

type AuthHandler struct {
	authSvc *service.AuthService
}

func NewAuthHandler(e *echo.Echo, authSvc *service.AuthService) {
	handler := &AuthHandler{
		authSvc: authSvc,
	}
	g := e.Group("/api/auth")
	g.POST("/register", handler.Register)
	g.POST("/login", handler.Login)
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req presenter.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.authSvc.Register(c.Request().Context(), req); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req presenter.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	accessToken, err := h.authSvc.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, presenter.LoginResponse{AccessToken: accessToken})
}
