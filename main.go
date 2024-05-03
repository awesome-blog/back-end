package main

import (
	"fmt"
	"pride/configs"
	"pride/database"
	"pride/modules/blog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"*"},
	}))

	database.ConnectDB()

	api_v1 := router.Group("/api/v1")
	{
		blog.Controller(api_v1)

	}

	router.Run(fmt.Sprintf("0.0.0.0:%s", configs.GetEnv("PORT")))
}
