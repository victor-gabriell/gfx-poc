package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("server",
	fx.Provide(NewHttpServer),
	fx.Invoke(func(engine *gin.Engine) {}),
)
