package http

import (
	"net/http"

	"github.com/cleidison-barradas/korvia.api/internal/services/application/dto"
	"github.com/cleidison-barradas/korvia.api/internal/services/application/usecases"
	"github.com/cleidison-barradas/korvia.api/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ServiceHandler struct {
	createUC *usecases.CreateServiceUseCase
	listUC   *usecases.ListServiceUseCase
	getUC    *usecases.GetServiceUseCase
	updateUC *usecases.UpdateServiceUseCase
}

func NewServiceHandler(
	createUC *usecases.CreateServiceUseCase, 
	listUC *usecases.ListServiceUseCase, 
	getUC *usecases.GetServiceUseCase, 
	updateUC *usecases.UpdateServiceUseCase,
	) *ServiceHandler {
	return &ServiceHandler{
		createUC: createUC,
		listUC:   listUC,
		getUC:    getUC,
		updateUC: updateUC,
	}
}

func (h *ServiceHandler) CreateServiceHandler(ctx *gin.Context) {
	 req := dto.CreateServiceRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Error(ctx, http.StatusBadRequest, utils.AppError{
			Code: utils.VALIDATION_ERROR,
			Message: err.Error(),
		})
		return
	}

	service, err := h.createUC.Execute(req)

	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, utils.AppError{
			Code: utils.CREATING_ERROR,
			Message: err.Error(),
		})
	}

	utils.Success(ctx, http.StatusCreated, service)
}

func (h *ServiceHandler) ListServiceHandler(ctx *gin.Context) {

	establishmentID := ctx.Param("establishmentId")
	
	parsedId, err := uuid.Parse(establishmentID)

	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, utils.AppError{
			Code: utils.VALIDATION_ERROR,
			Message: err.Error(),
		})
		return
	}

	services, err := h.listUC.Execute(parsedId)

	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, utils.AppError{
			Code: utils.LISTING_ERROR,
			Message: err.Error(),
		})
		return
	}

	utils.Success(ctx, http.StatusOK, services)
}

func (h *ServiceHandler) GetServiceHandler(ctx *gin.Context) {}
func (h *ServiceHandler) UpdateServiceHandler(ctx *gin.Context) {}

func (h *ServiceHandler) RegisterRoutes(router *gin.RouterGroup) {
	r := router.Group("/services")
	{
		r.POST("",h.CreateServiceHandler)
		r.GET("/:establishmentId",h.ListServiceHandler)
		r.PUT("/:serviceId",h.UpdateServiceHandler)
	}
}