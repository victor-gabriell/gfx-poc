package company

import (
	"github.com/victor-gabriell/gfx/internal/company/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("company",
	fx.Provide(
		repository.NewRepository,
		fx.Annotate(NewHandler, fx.From(new(*repository.CompanyRepository))),
	),
	fx.Invoke(SetRoutes),
)
