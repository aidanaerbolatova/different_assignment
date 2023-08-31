package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"rest/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckEmail(ctx *gin.Context) {
	var request models.CheckRequest
	err := json.NewDecoder(ctx.Request.Body).Decode(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(200, gin.H{"emails": FindEmails(request)})
}

func FindEmails(request models.CheckRequest) []string {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	regex := regexp.MustCompile(`Email:\s+(\S+)`)
	matches := regex.FindAllStringSubmatch(request.Body, -1)

	emails := make([]string, 0, len(matches))
	for _, match := range matches {
		if emailRegex.MatchString(match[1]) {
			emails = append(emails, match[1])
		}
	}
	return emails
}
