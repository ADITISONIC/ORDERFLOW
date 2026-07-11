package routes

import (
	"orderflow/handlers"
    "orderflow/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	authorized := router.Group("/")
    authorized.Use(middleware.AuthMiddleware())

    authorized.GET("/profile", handlers.Profile)
	authorized.POST("/orders", handlers.CreateOrder)
	authorized.GET("/orders", handlers.GetOrders)

	return router
}