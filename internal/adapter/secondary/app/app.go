package app

import (
	"context"
	"sync"

	"github.com/mehmetkmrc/nasilim.git/internal/adapter/secondary/config"
	"github.com/mehmetkmrc/nasilim.git/internal/core/port/cache"
	"github.com/mehmetkmrc/nasilim.git/internal/core/port/db"
	"github.com/mehmetkmrc/nasilim.git/internal/core/port/http"
	"go.uber.org/zap"
)

type App struct {
	rw *sync.RWMutex
	Cfg *config.Container
	HTTP http.ServerMaker
	PG db.EngineMaker
	MemCache cache.Memcache
	MemCacheTTL cache.MemcacheTTL
}

func New(
	rw *sync.RWMutex,
	cfg *config.Container,
	http http.ServerMaker,
	pg db.EngineMaker,
	memCache cache.Memcache,
	memCacheTTL cache.MemcacheTTL,
) *App{
	return &App{
		rw: rw,
		Cfg: cfg,
		HTTP: http,
		PG: pg,
		MemCache: memCache,
		MemCacheTTL: memCacheTTL,

	}
}
func (a *App) Run(ctx context.Context){
	zap.S().Info("RUNNER!")
}