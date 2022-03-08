package main

import (
	"github.com/gin-gonic/gin"
	"github.com/himanshuk42/gin-mongo-api/configs"
	"github.com/himanshuk42/gin-mongo-api/routes"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(router)

	router.Run("localhost:8090")
}
