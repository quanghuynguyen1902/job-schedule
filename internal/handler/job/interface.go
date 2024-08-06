package job

import "github.com/gin-gonic/gin"

type IHandler interface {
	CreateJob(c *gin.Context)
}
