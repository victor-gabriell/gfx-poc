package company

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine, handler *Handler) {
	r := g.Group("/company")

	r.Handle(http.MethodPost, "", handler.CreateCompany)
	r.Handle(http.MethodGet, ":id", handler.GetCompany)
}
