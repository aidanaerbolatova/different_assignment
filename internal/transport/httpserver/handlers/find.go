package handlers

import (
	"encoding/json"
	"net/http"
	"rest/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FindSubstr(ctx *gin.Context) {
	var request models.FindRequest
	err := json.NewDecoder(ctx.Request.Body).Decode(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(200, gin.H{"longestUniqueString": LongestUniqueStr(request)})
}

func LongestUniqueStr(request models.FindRequest) string {
	lastSeen := make(map[rune]int)
	start := 0
	maxLength := 0
	maxStart := 0

	for end, char := range request.Body {
		if lastIndex, exists := lastSeen[char]; exists && lastIndex >= start {
			start = lastIndex + 1
		}
		lastSeen[char] = end
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
			maxStart = start
		}
	}

	return request.Body[maxStart : maxStart+maxLength]
}
