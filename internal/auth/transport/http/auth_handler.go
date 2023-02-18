package transport

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/flashcard-api/internal/auth/presenter"
	"github.com/minhquy1903/flashcard-api/internal/auth/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req presenter.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.authService.Register(c.Request().Context(), req); err != nil {
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

	accessToken, err := h.authService.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, presenter.LoginResponse{AccessToken: accessToken})
}
