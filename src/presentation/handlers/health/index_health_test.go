package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/achmdndy/safa-life-api/src/application/health/dto"
)

// MockQueryHandler implements shared.QueryHandlerInterface for testing
type MockQueryHandler struct {
	mock.Mock
}

func (m *MockQueryHandler) Handle(query interface{}) (interface{}, error) {
	args := m.Called(query)
	return args.Get(0), args.Error(1)
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestHandler_GetHealth_Success(t *testing.T) {
	// Arrange
	router := setupTestRouter()
	mockHandler := new(MockQueryHandler)
	handler := NewHandler(mockHandler)
	
	// Mock the expected response
	expectedStatus := dto.HealthResponse{
		Status: "healthy",
		Services: map[string]string{
			"database": "ok",
			"redis":    "ok",
		},
	}
	mockHandler.On("Handle", mock.Anything).Return(expectedStatus, nil)
	
	router.GET("/health", handler.GetHealth)

	// Act
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response["success"].(bool))
	assert.Equal(t, "Health check completed", response["message"])
	assert.NotNil(t, response["data"])
	
	mockHandler.AssertExpectations(t)
}

func TestHandler_GetHealth_InternalError(t *testing.T) {
	// Arrange
	router := setupTestRouter()
	mockHandler := new(MockQueryHandler)
	handler := NewHandler(mockHandler)
	
	// Mock an error response
	mockHandler.On("Handle", mock.Anything).Return(nil, assert.AnError)
	
	router.GET("/health", handler.GetHealth)

	// Act
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["success"].(bool))
	
	mockHandler.AssertExpectations(t)
}

func TestHandler_GetHealthSimple_Success(t *testing.T) {
	// Arrange
	router := setupTestRouter()
	mockHandler := new(MockQueryHandler)
	handler := NewHandler(mockHandler)
	
	router.GET("/health/simple", handler.GetHealthSimple)

	// Act
	req, _ := http.NewRequest("GET", "/health/simple", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response["success"].(bool))
	assert.Equal(t, "Simple health check", response["message"])
	assert.NotNil(t, response["data"])
	
	// Check the data structure
	data, ok := response["data"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "ok", data["status"])
	assert.NotNil(t, data["time"])
}

func TestNewHandler(t *testing.T) {
	// Arrange
	mockHandler := new(MockQueryHandler)

	// Act
	handler := NewHandler(mockHandler)

	// Assert
	assert.NotNil(t, handler)
	assert.Equal(t, mockHandler, handler.healthQueryHandler)
}