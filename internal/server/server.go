package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHttpServer(lc fx.Lifecycle) *gin.Engine {
	engine := gin.Default()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go engine.Run(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return engine
}
