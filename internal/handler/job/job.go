package job

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quanghuynguyen1902/job-schedule/internal/config"
	"github.com/quanghuynguyen1902/job-schedule/internal/controller"
	"github.com/quanghuynguyen1902/job-schedule/internal/db"
	"github.com/quanghuynguyen1902/job-schedule/internal/logger"
	"github.com/quanghuynguyen1902/job-schedule/internal/model"
	"github.com/quanghuynguyen1902/job-schedule/internal/view"
)

type handler struct {
	controller *controller.Controller
	store      *db.Store
	logger     logger.Logger
	config     *config.Config
}

// New returns a handler
func New(store *db.Store, ctrl *controller.Controller, logger logger.Logger, cfg *config.Config) IHandler {
	return &handler{
		controller: ctrl,
		store:      store,
		logger:     logger,
		config:     cfg,
	}
}

func (h handler) CreateJob(c *gin.Context) {
	l := h.logger.Fields(logger.Fields{
		"handler": "job",
		"method":  "CreateJob",
	})

	var job model.Job

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, view.CreateResponse[any](nil, nil, err, job, ""))
		return
	}

	res, err := h.controller.Job.CreateJob(&job)
	if err != nil || res == nil {
		l.Error(err, "failed to create job")
		c.JSON(http.StatusInternalServerError, view.CreateResponse[any](nil, nil, err, nil, ""))
		return
	}

	c.JSON(http.StatusOK, view.CreateResponse[any](view.ToJob(*res), nil, nil, nil, ""))
}
