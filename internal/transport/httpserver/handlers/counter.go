package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CounterAdd(ctx *gin.Context) {
	i := ctx.Param("i")
	convertI, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.service.Cache.AddCounter(ctx, "counter", int64(convertI)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}

func (h *Handler) CounterSub(ctx *gin.Context) {
	i := ctx.Param("i")
	convertI, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.service.Cache.SubCounter(ctx, "counter", int64(convertI)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}

func (h *Handler) CounterVal(ctx *gin.Context) {
	counter, err := h.service.Cache.GetCounter(ctx, "counter")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(200, gin.H{"Counter Value": counter})
}
