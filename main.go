package main

import (
	"fmt"
	"ingestion-service/handler"
	"ingestion-service/kafka"
	"ingestion-service/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("env file not found, continuing with system env")
	}

	// ✅ initialize Kafka producer
	kafka.InitKafkaProducer()

	r := gin.Default()

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/message", handler.MessageHandler)
	}

	r.Run(":" + os.Getenv("PORT"))
}
