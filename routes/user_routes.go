package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/himanshuk42/gin-mongo-api/controllers"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controllers.GetAllUsers())
	router.GET("/user/:userId", controllers.GetAUser())
	router.POST("/user", controllers.CreateUser())
	router.PUT("/user/:userId", controllers.EditUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
}
