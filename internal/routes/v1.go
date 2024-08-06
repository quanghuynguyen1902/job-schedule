package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/quanghuynguyen1902/job-schedule/internal/handler"
)

func loadV1Routes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1")
	job := api.Group("/job")
	{
		job.POST("/", h.Job.CreateJob)
	}
}
