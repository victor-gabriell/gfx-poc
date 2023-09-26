package repository

import (
	"github.com/google/uuid"
	"github.com/victor-gabriell/gfx/internal/company/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateCompany(company model.Company) (model.Company, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return model.Company{}, err
	}

	company.Id = id

	err = r.db.Create(&company).Error

	return company, err
}

func (r *Repository) GetCompanyById(id uuid.UUID) (*model.Company, error) {
	var company model.Company
	if err := r.db.First(&company, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &company, nil
}
