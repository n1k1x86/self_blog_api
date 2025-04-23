package server

import (
	"api/config"
	handler "api/internal/services/server/handlers"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GinServer struct {
	Engine   *gin.Engine
	Pool     *pgxpool.Pool
	CancelFn context.CancelFunc
}

func CreateNewServer(cfg config.BlogDBConfig) (*GinServer, error) {
	ctx, cancel := context.WithCancel(context.Background())

	config, err := pgxpool.ParseConfig(cfg.BuildDSN())
	if err != nil {
		cancel()
		return nil, err
	}
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		cancel()
		return nil, err
	}

	router := gin.Default()

	handler.SetHandlersTags(router, pool)

	return &GinServer{
		Engine:   router,
		Pool:     pool,
		CancelFn: cancel,
	}, nil
}
