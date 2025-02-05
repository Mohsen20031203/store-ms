package api

import (
	"hello/config"
	"hello/internal/db"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Store  *gorm.DB
	Config config.Config
	Router *gin.Engine
}

func NewServer(storege *db.Storege, config *config.Config) (*Server, error) {

	server := &Server{
		Store:  storege.DB,
		Config: *config,
	}
	server.setupRouter()
	return server, nil

}

func (s *Server) setupRouter() {
	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"massage": "Ok",
		})
	})
	s.Router = router
}
