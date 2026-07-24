package http

import (
	"net/http"

	"github.com/cleidison-barradas/korvia.api/internal/users/application/dto"
	"github.com/cleidison-barradas/korvia.api/internal/users/application/usecases"
	"github.com/cleidison-barradas/korvia.api/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUserUC *usecases.CreateUserUseCase
	updateUserUC *usecases.UpdateUserUseCase
	activateUserUC *usecases.ActivateUserUseCase
}

func NewUserHandler(
	createUserUC *usecases.CreateUserUseCase, 
	updateUserUC *usecases.UpdateUserUseCase, 
	activateUserUC *usecases.ActivateUserUseCase) *UserHandler {
	return &UserHandler{
		createUserUC: createUserUC,
		updateUserUC: updateUserUC,
		activateUserUC: activateUserUC,
	}
}

func (h *UserHandler) CreateUserHandler(ctx *gin.Context) {

	req := dto.CreateUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Error(ctx, http.StatusBadRequest, utils.AppError{
			Code: "payload_error",
			Message: err.Error(),
		})
		return
	}

	user, err := h.createUserUC.Execute(req)

	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, utils.AppError{
			Code: "error_creating_user",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)

}
func (u *UserHandler) UpdateUserHandler(ctx *gin.Context) {}
func (u *UserHandler) ActivateUserHandler(ctx *gin.Context) {}

func (h *UserHandler) RegisterRoutes(router *gin.RouterGroup) {
	r := router.Group("/users")
	{
		r.POST("",h.CreateUserHandler)
		r.PUT("/:userId", h.UpdateUserHandler)
		r.PUT("/:userId/activate", h.ActivateUserHandler)
	}
}