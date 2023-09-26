package company

import (
	"github.com/victor-gabriell/gfx/internal/company/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("company",
	fx.Provide(
		NewHandler,
		repository.NewRepository,
	),
	fx.Invoke(SetRoutes),
)
