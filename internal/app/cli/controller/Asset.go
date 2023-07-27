package controller

import (
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ka1i/cli/internal/pkg/response"
	"github.com/ka1i/cli/pkg/assets"
)

func EmbedAsset(ctx *gin.Context) {
	var storage = assets.GetStorage()
	file := ctx.Param("filepath")
	log.Println(file, ctx.Writer.Status())
	if file == "/" || file == "/index.html" || file == "/gin.txt" {
		ctx.Status(http.StatusOK)

		ctx.FileFromFS(filepath.Join(storage.Root, "/gin.txt"), http.FS(storage.Fs))
		return
	}

	response.Fail(ctx, gin.H{
		"file":      file,
		"status":    http.StatusNotFound,
		"timestamp": time.Now().UnixMilli(),
	}, "file server serve")
}
