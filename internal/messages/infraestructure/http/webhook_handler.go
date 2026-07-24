package http

import (
	"net/http"

	"github.com/cleidison-barradas/korvia.api/internal/messages/application/dto"
	"github.com/cleidison-barradas/korvia.api/internal/messages/application/usecases"
	"github.com/cleidison-barradas/korvia.api/internal/utils"
	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	uc *usecases.IncomingMessageUseCase
}

func NewWebhookHandler(usecase *usecases.IncomingMessageUseCase) *WebhookHandler {
	return &WebhookHandler{
		uc: usecase,
	}
}

func (h *WebhookHandler) IncomingMessageHandler(ctx *gin.Context) {
	var req dto.WebhookPayload

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Error(ctx, http.StatusBadRequest, utils.AppError{
			Code: utils.VALIDATION_ERROR,
			Message: err.Error(),
		})
		return
	}

	var wabaID = req.Entry[0].ID
	var phoneNumber = req.Entry[0].Changes[0].Value.Messages[0].From
	var message = req.Entry[0].Changes[0].Value.Messages[0].Text.Body
	var name = req.Entry[0].Changes[0].Value.Contacts[0].Profile.Name

	h.uc.Execute(dto.IncomingMessage{
		Name: name,
		WabaID: wabaID,
		Message: message,
		PhoneNumber: phoneNumber,
	})
}

func (h *WebhookHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/webhook", h.IncomingMessageHandler)
}