package handler

import (
	"L0/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	order := router.Group("/order")
	{
		order.POST("/", h.createOrder)
		order.GET("/id", h.getOrder)
	}

	return router
}
