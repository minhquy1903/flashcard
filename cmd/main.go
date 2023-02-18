package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/flashcard-api/config"
	"github.com/minhquy1903/flashcard-api/db"
	"github.com/minhquy1903/flashcard-api/internal/auth/repository"
	"github.com/minhquy1903/flashcard-api/internal/auth/service"
	httpTransport "github.com/minhquy1903/flashcard-api/internal/auth/transport/http"
	"github.com/minhquy1903/flashcard-api/pkg/token"
	"github.com/minhquy1903/flashcard-api/server"
)

func main() {

	// Load config
	cfg, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	// Connect and get DB instance
	db := db.GetPostgresInstance(cfg)

	tokenManager := token.NewTokenManager(cfg.JWTSecret)

	// Register all repositories
	userRepo := repository.NewUserRepository(db)

	// Register all services
	authSvc := service.NewAuthService(userRepo, tokenManager)

	// Register all handlers
	authHandler := httpTransport.NewAuthHandler(authSvc)

	// New server
	svr := server.NewServer(cfg, echo.New())

	authGroup := svr.Echo.Group("auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)
	svr.Echo.GET("/heath", func(c echo.Context) error { return c.JSON(http.StatusOK, "alive") })

	svr.Run()
}
