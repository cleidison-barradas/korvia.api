package router

import "github.com/gin-gonic/gin"

type Register interface {
	RegisterRoutes(router *gin.RouterGroup)
}