package handler

import (
	"encoding/json"
	"ingestion-service/kafka"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	var req struct {
		Message string `json:"message"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid content"})
		return
	}

	message := map[string]interface{}{
		"user_id": userID,
		"message": req.Message,
	}

	payload, _ := json.Marshal(message)

	log.Println(string(payload))

	if err := kafka.PublishMessage(payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to publish", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "message accepted"})
}
