package routes

import (
	"orderflow/handlers"
	"orderflow/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

    ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.RateLimiter())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	authorized.GET("/profile", handlers.Profile)
	authorized.POST("/orders",middleware.IdempotencyMiddleware(), handlers.CreateOrder)
	authorized.GET("/orders", handlers.GetOrders)
	authorized.GET("/metrics", handlers.GetMetrics)

	return router
}
