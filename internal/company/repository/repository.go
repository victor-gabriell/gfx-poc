package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/victor-gabriell/gfx/internal/company/model"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (r *CompanyRepository) CreateCompany(ctx context.Context, company model.Company) (model.Company, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return model.Company{}, err
	}

	company.Id = id

	err = r.db.WithContext(ctx).Create(&company).Error

	return company, err
}

func (r *CompanyRepository) GetCompanyById(ctx context.Context, id uuid.UUID) (*model.Company, error) {
	var company model.Company
	if err := r.db.WithContext(ctx).First(&company, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &company, nil
}
