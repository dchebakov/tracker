package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dchebakov/tracker/config"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Server struct {
	echo     *echo.Echo
	cfg      *config.Config
	db       *sqlx.DB
	logger   *zap.SugaredLogger
	validate *validator.Validate
}

func NewServer(
	cfg *config.Config,
	db *sqlx.DB,
	logger *zap.SugaredLogger,
	validate *validator.Validate,
) *Server {
	return &Server{echo: echo.New(), cfg: cfg, db: db, logger: logger, validate: validate}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr: s.cfg.Api.Host + ":" + s.cfg.Api.Port,
	}

	go func() {
		err := s.echo.StartServer(server)
		if err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}

		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Api.Port)
	}()

	s.RegisterHandlers()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
