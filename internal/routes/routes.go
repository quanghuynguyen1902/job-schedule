package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/quanghuynguyen1902/job-schedule/internal/db"

	"github.com/quanghuynguyen1902/job-schedule/internal/config"
	"github.com/quanghuynguyen1902/job-schedule/internal/controller"
	"github.com/quanghuynguyen1902/job-schedule/internal/handler"
	"github.com/quanghuynguyen1902/job-schedule/internal/logger"
)

func setupCORS(r *gin.Engine, cfg *config.Config) {
	r.Use(func(c *gin.Context) {
		cors.New(
			cors.Config{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders: []string{
					"Origin", "Host", "Content-Type", "Content-Length", "Accept-Encoding", "Accept-Language", "Accept",
					"X-CSRF-Token", "Authorization", "X-Requested-With", "X-Access-Token",
				},
				AllowCredentials: true,
			},
		)(c)
	})
}

func NewRoutes(cfg *config.Config, s *db.Store, repo db.DBRepo, logger logger.Logger) *gin.Engine {
	r := gin.New()
	pprof.Register(r)

	ctrl := controller.New(s, repo, logger, cfg)
	h := handler.New(s, ctrl, logger, cfg)

	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/healthz"),
		gin.Recovery(),
	)
	// config CORS
	setupCORS(r, cfg)

	// load API here
	loadV1Routes(r, h)

	return r
}
