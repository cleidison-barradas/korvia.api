package http

import (
	"net/http"

	"github.com/cleidison-barradas/korvia.api/internal/establishments/application/dto"
	"github.com/cleidison-barradas/korvia.api/internal/establishments/application/usecases"
	"github.com/cleidison-barradas/korvia.api/internal/utils"
	"github.com/gin-gonic/gin"
)

type EstablishmentHandler struct {
	CreateUC *usecases.CreateEstablishmentUseCase
	UpdateUC *usecases.UpdateEstablishmentUseCase
}

func NewEstablishmentHandler(p *EstablishmentHandler) *EstablishmentHandler {
	return &EstablishmentHandler{
		CreateUC: p.CreateUC,
		UpdateUC: p.UpdateUC,
	}
}

func (h *EstablishmentHandler) CreateEstablishmentHandler(ctx *gin.Context) {

	req := dto.CreateEstablishmentRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Error(ctx, http.StatusBadRequest, utils.AppError{
			Code: "payload_error",
			Message: err.Error(),
		})
		return
	}

	establishment, err := h.CreateUC.Execute(req)

	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, utils.AppError{
			Code: "error_creating_establishment",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, establishment)
}

func (h *EstablishmentHandler) UpdateHandler(ctx *gin.Context) {
	
}

func (h *EstablishmentHandler) RegisterRoutes(router *gin.RouterGroup) {
	r := router.Group("/establishments")
	{
		r.POST("",h.CreateEstablishmentHandler)
		r.PUT("/:establishmentId", h.UpdateHandler)
	}	
}