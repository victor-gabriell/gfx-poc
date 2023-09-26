package company

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/victor-gabriell/gfx/internal/company/model"
	"github.com/victor-gabriell/gfx/internal/company/repository"
)

type Handler struct {
	companyRepo repository.Repository
}

func NewHandler(companyRepo repository.Repository) *Handler {
	return &Handler{
		companyRepo: companyRepo,
	}
}

func (h *Handler) CreateCompany(g *gin.Context) {
	var company model.Company

	if err := g.ShouldBindWith(&company, binding.JSON); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	createCompany, err := h.companyRepo.CreateCompany(g, company)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusCreated, createCompany)
}

func (h *Handler) GetCompany(g *gin.Context) {
	id := g.Param("id")

	parse, err := uuid.Parse(id)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	company, err := h.companyRepo.GetCompanyById(g, parse)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusOK, company)
}
