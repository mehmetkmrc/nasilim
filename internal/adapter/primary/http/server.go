package http

import (
	"context"
	"github.com/goccy/go-json"
	"errors"
	"fmt"
	std_http "net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/mehmetkmrc/nasilim.git/internal/core/port/http"
	"github.com/mehmetkmrc/nasilim.git/internal/adapter/secondary/config"
	"go.uber.org/zap"
)

const (
	viewPath = "../../internal/adapter/primary/http/web/views"
	//../../internal/adapter/transport/http/web/views
	publicPath = "../../internal/adapter/primary/http/web/public"
	renderType = ".html"
	assetPath  = "./asset"
)

var (
	_ http.ServerMaker = (*server)(nil)
)

type(
	server struct{
		ctx 		context.Context
		cfg			*config.Container
		app 		*fiber.App
		cfgFiber	*fiber.Config
	}
)

func NewHTTPServer(
	ctx context.Context,
	cfg *config.Container,
) http.ServerMaker {
	return &server{
		ctx: ctx,
		cfg: cfg,
	}
}

func (s *server) Start(ctx context.Context) error{
	engine := html.New(viewPath, renderType)

	app := fiber.New(fiber.Config{
		ReadTimeout: time.Minute*time.Duration(s.cfg.Settings.ServerReadTimeout),
		StrictRouting: false,
		CaseSensitive: true,
		BodyLimit: 4*1024*1024,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		AppName: "Go-Gateway",
		Immutable: true,
		Views: engine,
		ViewsLayout: "layouts/main",
	})
	// app.Static("/", publicPath, fiber.Static{
	//  	Index: "main.html",
	// })
	s.app = app
	fiberConnURL := fmt.Sprintf("%s:%d", s.cfg.HTTP.Host, s.cfg.HTTP.Port)

	go func ()  {
		zap.S().Info("Starting HTTP server on", fiberConnURL)
		if err := s.app.Listen(fiberConnURL); err != nil{
			if errors.Is(err, std_http.ErrServerClosed){
				return 
			}
			zap.S().Fatal("server listen error: %w", err)
		}
	}()
	err := s.HTTPMiddleware()
	if err != nil{
		zap.S().Fatal("middleware error:", err)
	}
	s.SetupRoutes()

	return nil
}

func (s *server) Close(ctx context.Context) error{
	zap.S().Info("HTTP-Server Context is done. Shutting down server...")
	if err := s.app.ShutdownWithContext(ctx); err != nil{
		zap.S().Info("server shutdown error: %w", err)
		return err
	}
	return nil
}