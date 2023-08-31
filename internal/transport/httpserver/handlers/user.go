package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"rest/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrorNoRows = errors.New("user with such ID does not exists")

func (h *Handler) CreateUser(ctx *gin.Context) {
	var user models.User

	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := h.service.User.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"user": user, "result": "User created successfully"})
}

func (h *Handler) GetUser(ctx *gin.Context) {
	var user models.User

	id := ctx.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	user.Id = convertId

	result, err := h.service.User.GetUser(user)
	if err != nil {
		if errors.Is(err, ErrorNoRows) {
			ctx.AbortWithError(http.StatusBadRequest, ErrorNoRows)
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"user": result, "result": "Get user successfully"})
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	var user models.User

	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	id := ctx.Param("id")
	convertID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	user.Id = convertID

	if err = h.service.User.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"result": "User updated successfully"})
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	var user models.User

	id := ctx.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	user.Id = convertId

	if err = h.service.User.DeleteUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"result": "User deleted successfully"})
}
