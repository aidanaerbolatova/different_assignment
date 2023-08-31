package handlers

import (
	"rest/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	service *service.Service
	logger  *zap.SugaredLogger
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoute() *gin.Engine {
	router := gin.Default()

	router.POST("/rest/substr/find", h.FindSubstr)

	router.POST("/rest/email/check", h.CheckEmail)

	router.POST("/rest/counter/add/:i", h.CounterAdd)
	router.POST("/rest/counter/sub/:i", h.CounterSub)
	router.GET("/rest/counter/val", h.CounterVal)

	router.POST("/rest/user", h.CreateUser)
	router.GET("/rest/user/:id", h.GetUser)
	router.PUT("/rest/user/:id", h.UpdateUser)
	router.DELETE("/rest/user/:id", h.DeleteUser)

	return router
}
