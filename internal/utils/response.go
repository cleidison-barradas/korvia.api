package utils

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Data any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
}

type AppError struct {
	Code ERROR_CODE `json:"code"`
	Message string `json:"message"`
}

func Success(ctx *gin.Context, statusCode int, data any) {
	ctx.JSON(statusCode, ApiResponse{
		Data: data,
	})
}

func Error(ctx *gin.Context, statusCode int, err AppError) {
	ctx.JSON(statusCode, ApiResponse{
		Error: err,
	})
}

func MiddlewareError(ctx *gin.Context, statusCode int, err AppError) {
	ctx.AbortWithStatusJSON(statusCode, ApiResponse{
		Error: err,
	})
}