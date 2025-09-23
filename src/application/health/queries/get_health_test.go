package queries

import (
	"testing"

	"github.com/achmdndy/safa-life-api/src/application/health/dto"
	"github.com/achmdndy/safa-life-api/src/domain/health"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHealthService is a mock implementation of health.HealthServiceInterface
type MockHealthService struct {
	mock.Mock
}

func (m *MockHealthService) GetHealth() health.HealthStatus {
	args := m.Called()
	return args.Get(0).(health.HealthStatus)
}

func TestGetHealthQueryHandler_Handle_Success(t *testing.T) {
	// Arrange
	mockService := new(MockHealthService)
	expectedStatus := health.HealthStatus{
		Status: "ok",
		Services: map[string]string{
			"database": "ok",
			"redis":    "ok",
		},
	}
	mockService.On("GetHealth").Return(expectedStatus)

	handler := NewGetHealthQueryHandler(mockService)
	query := GetHealthQuery{}

	// Act
	result, err := handler.Handle(query)

	// Assert
	assert.NoError(t, err)
	
	// Type assert to DTO
	dtoResult, ok := result.(dto.HealthResponse)
	assert.True(t, ok)
	assert.Equal(t, "ok", dtoResult.Status)
	assert.Equal(t, "ok", dtoResult.Services["database"])
	assert.Equal(t, "ok", dtoResult.Services["redis"])
	
	mockService.AssertExpectations(t)
}

func TestGetHealthQueryHandler_Handle_DegradedStatus(t *testing.T) {
	// Arrange
	mockService := new(MockHealthService)
	expectedStatus := health.HealthStatus{
		Status: "degraded",
		Services: map[string]string{
			"database": "ok",
			"redis":    "error",
		},
	}
	mockService.On("GetHealth").Return(expectedStatus)

	handler := NewGetHealthQueryHandler(mockService)
	query := GetHealthQuery{}

	// Act
	result, err := handler.Handle(query)

	// Assert
	assert.NoError(t, err)
	
	// Type assert to DTO
	dtoResult, ok := result.(dto.HealthResponse)
	assert.True(t, ok)
	assert.Equal(t, "degraded", dtoResult.Status)
	assert.Equal(t, "ok", dtoResult.Services["database"])
	assert.Equal(t, "error", dtoResult.Services["redis"])
	
	mockService.AssertExpectations(t)
}

func TestNewGetHealthQueryHandler(t *testing.T) {
	// Arrange
	mockService := new(MockHealthService)

	// Act
	handler := NewGetHealthQueryHandler(mockService)

	// Assert
	assert.NotNil(t, handler)
	assert.IsType(t, &GetHealthQueryHandler{}, handler)
}
