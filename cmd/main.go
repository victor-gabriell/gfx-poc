package main

import (
	"github.com/victor-gabriell/gfx/internal/company"
	"github.com/victor-gabriell/gfx/internal/config"
	"github.com/victor-gabriell/gfx/internal/db"
	"github.com/victor-gabriell/gfx/internal/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.NewApplicationConfig),
		server.Module,
		db.Module,
		company.Module,
	).Run()
}
