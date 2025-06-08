package test

import (
	"bytes"
	"encoding/json"
	"ingestion-service/handler"
	"ingestion-service/kafka"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Set Kafka.Writer ke mock sebelum semua test
	kafka.Writer = &kafka.MockWriter{}
	os.Exit(m.Run())
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/message", handler.MessageHandler)
	return r
}

func TestMessageHandler_Success(t *testing.T) {
	router := setupRouter()

	payload := map[string]interface{}{
		"user_id": 1,
		"message": "Hello from test",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/message", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "success")
}

func TestMessageHandler_InvalidJSON(t *testing.T) {
	router := setupRouter()

	body := []byte(`invalid-json`)
	req, _ := http.NewRequest("POST", "/api/message", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid JSON")
}
