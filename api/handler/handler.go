package handler

import (
	"exam3/api/models"
	"exam3/pkg/logger"
	"exam3/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.IServiceManager
	log     logger.ILogger
}

func New(manager service.IServiceManager, log logger.ILogger) *Handler {
	return &Handler{
		service: manager,
		log:     log,
	}
}

func handleResponse(c *gin.Context, log logger.ILogger, msg string, statusCode int, data interface{}) {
	response := models.Response{}

	switch code := statusCode; {
	case code < 400:
		response.Description = "OK"
		log.Info("OK", logger.String("msg", msg), logger.Int("status", statusCode))
	case code < 500:
		response.Description = "Bad Request"
		log.Info("BAD REQUEST", logger.String("msg", msg), logger.Int("status", statusCode))
	default:
		response.Description = "Internal Server Error"
		log.Info("INTERNAL SERVER ERROR", logger.String("msg", msg), logger.Int("status", statusCode))
	}

	response.Data = data
	response.StatusCode = statusCode

	c.JSON(response.StatusCode, response)
}
