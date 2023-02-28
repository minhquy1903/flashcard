package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/flashcard-api/config"
	"github.com/minhquy1903/flashcard-api/db"
	ar "github.com/minhquy1903/flashcard-api/internal/auth/repository"
	av "github.com/minhquy1903/flashcard-api/internal/auth/service"
	ah "github.com/minhquy1903/flashcard-api/internal/auth/transport/http"
	vr "github.com/minhquy1903/flashcard-api/internal/vocabulary/repository"
	vs "github.com/minhquy1903/flashcard-api/internal/vocabulary/service"
	vh "github.com/minhquy1903/flashcard-api/internal/vocabulary/transport/http"
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

	// New token manager
	tm := token.NewTokenManager(cfg.JWTSecret)

	// Register all repositories
	ur := ar.NewUserRepository(db)
	vr := vr.NewVocabRepository(db)

	// Register all services
	as := av.NewAuthService(ur, tm)
	vs := vs.NewVocabularyService(vr)

	// New server
	svr := server.NewServer(cfg, echo.New())

	// Register all handlers
	ah.NewAuthHandler(svr.Echo, as)
	vh.NewVocabularyHandler(svr.Echo, vs)

	svr.Echo.GET("/heath", func(c echo.Context) error { return c.JSON(http.StatusOK, "alive") })

	svr.Run()
}
