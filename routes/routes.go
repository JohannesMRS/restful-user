package routes

import (
	"backend-api/controller"
	"backend-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.POST("/api/register", controller.Register)
	router.POST("/api/login", controller.Login)
	router.GET("/api/users", middlewares.AuthMiddleware(), controller.FindUsers)
	router.POST("/api/users", middlewares.AuthMiddleware(), controller.CreateUser)
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controller.FindUserById)
	router.PUT("api/users/:id", middlewares.AuthMiddleware(), controller.UpdateUser)
	router.DELETE("api/users/:id", middlewares.AuthMiddleware(), controller.DeleteUser)
	return router
}
