package api

import (
	_ "exam3/api/docs"
	"exam3/api/handler"
	"exam3/pkg/logger"
	"exam3/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"time"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(manager service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.New(manager, log)

	r := gin.New()

	// middleware
	r.Use(traceRequest)

	{
		r.POST("/book", h.CreateBook)
		r.GET("/book/:id", h.GetBook)
		r.GET("/books", h.GetBookList)
		r.PUT("/book/:id", h.UpdateBook)
		r.DELETE("/book/:id", h.DeleteBook)
		r.PATCH("/book/:id", h.UpdateBookPage)

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	}

	return r
}

func traceRequest(c *gin.Context) {
	startTime := time.Now()
	duration := time.Since(startTime).Nanoseconds()

	log.Println("END TIME", time.Now().Format("15:04:05.000"), c.Request.Method, c.Request.URL.Path, "status", c.Writer.Status(), "time", duration)
}
