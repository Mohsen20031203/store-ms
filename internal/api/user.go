package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) User(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ok": "get user is true",
	})

}
