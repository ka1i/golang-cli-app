package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ka1i/cli/internal/pkg/response"
)

func EntryQuery(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"total":     0,
		"status":    http.StatusOK,
		"timestamp": time.Now().UnixMilli(),
		"data":      []string{},
	}, "entry query success")
}
