package server

import (
	blrepository "github.com/dchebakov/tracker/internal/blacklist/repository"
	blusecase "github.com/dchebakov/tracker/internal/blacklist/usecase"
	clhandler "github.com/dchebakov/tracker/internal/collector/handler"
	cusecase "github.com/dchebakov/tracker/internal/collector/usecase"
	ctrepository "github.com/dchebakov/tracker/internal/customer/repository"
	ctusecase "github.com/dchebakov/tracker/internal/customer/usecase"
	hlhandler "github.com/dchebakov/tracker/internal/health/handler"
	hlrepository "github.com/dchebakov/tracker/internal/health/repository"
	hlusecase "github.com/dchebakov/tracker/internal/health/usecase"
	tmiddlewares "github.com/dchebakov/tracker/internal/middleware"
	strepository "github.com/dchebakov/tracker/internal/stats/repository"
	stusecase "github.com/dchebakov/tracker/internal/stats/usecase"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterHandlers() {
	healthRepo := hlrepository.NewHealthRepository(s.db)
	customerRepo := ctrepository.NewCustomerRepository(s.db)
	blacklistRepo := blrepository.NewBlacklistRepository(s.db)
	statsRepo := strepository.NewStatsRepository(s.db, s.logger)

	healthUC := hlusecase.NewHealthUseCase(healthRepo)
	customerUC := ctusecase.NewCustomerUseCase(s.logger, customerRepo)
	blacklistUC := blusecase.NewBlacklistUseCase(s.logger, blacklistRepo)
	statsUC := stusecase.NewStatsUseCase(s.logger, statsRepo)
	collectorUC := cusecase.NewCollectorUseCase(s.logger, customerUC, blacklistUC, statsUC)

	healthHandler := hlhandler.NewHelthHandler(s.logger, healthUC)
	collectorHandler := clhandler.NewCollectorHandler(s.logger, s.validate, collectorUC)

	mw := tmiddlewares.NewMiddlewareManager(s.logger)
	s.echo.Use(middleware.BodyLimit("2M"))
	s.echo.Use(mw.RequestLoggerMiddleware)

	v1 := s.echo.Group("/api/v1")

	healthGroup := v1.Group("/health")
	healthHandler.RegisterHTTPEndPoints(healthGroup)

	collectorGroup := v1.Group("/collect")
	collectorHandler.RegisterHTTPEndPoints(collectorGroup)
}
