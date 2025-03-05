package api

import (
	"hello/internal/db/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreatCategory(ctx *gin.Context) {
	var category models.Category

	if err := ctx.BindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	err := s.Store.CreateCategory(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}
	ctx.JSON(http.StatusCreated, category)
}

func (s *Server) GetCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	idCategory, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is not valid"})
		return
	}

	category, err := s.Store.GetCategory(idCategory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}
	ctx.JSON(http.StatusCreated, category)

}

func (s *Server) ListCategory(ctx *gin.Context) {

	category, err := s.Store.ListCategory()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}
	ctx.JSON(http.StatusCreated, category)

}

func (s *Server) UpdateCategory(ctx *gin.Context) {
	var category models.Category

	err := ctx.BindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "internal server error",
		})
	}
}
