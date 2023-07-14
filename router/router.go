package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Controller contains the controller methods
type Controller interface {
	Ping(*gin.Context)
	GetSample(*gin.Context)
	GetAnswer(*gin.Context)
}

// New creates the router
func New(controller Controller) *gin.Engine {
	r := gin.New()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/v1")
	v1.GET("/ping", controller.Ping)
	v1.GET("/sample", controller.GetSample)

	v1.POST("/get-answer", controller.GetAnswer)
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
