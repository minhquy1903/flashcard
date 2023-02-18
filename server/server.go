package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minhquy1903/flashcard-api/config"
)

type Server struct {
	Cfg  *config.Config
	Echo *echo.Echo
}

func NewServer(cfg *config.Config, echo *echo.Echo) *Server {
	return &Server{Cfg: cfg, Echo: echo}
}

func (s *Server) Run() {
	s.Echo.Use(middleware.CORS())

	var ready chan bool

	server := &http.Server{
		Addr:         ":" + s.Cfg.ServerPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		fmt.Printf("Server is listening on PORT: %s", s.Cfg.ServerPort)

		if err := s.Echo.StartServer(server); err != nil {
			log.Fatalln("Error starting Server: ", err)
		}

	}()

	if ready != nil {
		ready <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Fatalln("Server Exited Properly")
	s.Echo.Shutdown(ctx)

}
