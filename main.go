package main

import (
	"log"

	"recruitment_task/config"
	"recruitment_task/data"
	"recruitment_task/handlers"
	"recruitment_task/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	configData, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	logger.SetLogLevel(configData.LogLevel)

	log.Printf("Info: Server is starting on port %s", configData.Port)

	dataList, err := data.LoadDataFromFile("input.txt")
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
		return
	}

	r := gin.Default()
	r.GET("/endpoint/:value", func(c *gin.Context) {
		handlers.EndpointHandler(c, dataList)
	})

	log.Printf("Info: Server is running on :%s", configData.Port)
	r.Run(":" + configData.Port)
}
