package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/victor-gabriell/gfx/internal/company/model"
)

type Repository interface {
	CreateCompany(ctx context.Context, company model.Company) (model.Company, error)
	GetCompanyById(ctx context.Context, id uuid.UUID) (*model.Company, error)
}
