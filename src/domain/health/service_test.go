package health

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCheckerRepository is a mock implementation of CheckerRepository
type MockCheckerRepository struct {
	mock.Mock
}

func (m *MockCheckerRepository) CheckDatabase() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCheckerRepository) CheckRedis() error {
	args := m.Called()
	return args.Error(0)
}

func TestService_GetHealth_AllServicesHealthy(t *testing.T) {
	// Arrange
	mockRepo := new(MockCheckerRepository)
	mockRepo.On("CheckDatabase").Return(nil)
	mockRepo.On("CheckRedis").Return(nil)
	
	service := NewService(mockRepo)

	// Act
	result := service.GetHealth()

	// Assert
	assert.Equal(t, "ok", result.Status)
	assert.Equal(t, "ok", result.Services["database"])
	assert.Equal(t, "ok", result.Services["redis"])
	mockRepo.AssertExpectations(t)
}

func TestService_GetHealth_DatabaseDown(t *testing.T) {
	// Arrange
	mockRepo := new(MockCheckerRepository)
	mockRepo.On("CheckDatabase").Return(errors.New("database connection failed"))
	mockRepo.On("CheckRedis").Return(nil)
	
	service := NewService(mockRepo)

	// Act
	result := service.GetHealth()

	// Assert
	assert.Equal(t, "degraded", result.Status)
	assert.Equal(t, "down", result.Services["database"])
	assert.Equal(t, "ok", result.Services["redis"])
	mockRepo.AssertExpectations(t)
}

func TestService_GetHealth_RedisDown(t *testing.T) {
	// Arrange
	mockRepo := new(MockCheckerRepository)
	mockRepo.On("CheckDatabase").Return(nil)
	mockRepo.On("CheckRedis").Return(errors.New("redis connection failed"))
	
	service := NewService(mockRepo)

	// Act
	result := service.GetHealth()

	// Assert
	assert.Equal(t, "degraded", result.Status)
	assert.Equal(t, "ok", result.Services["database"])
	assert.Equal(t, "down", result.Services["redis"])
	mockRepo.AssertExpectations(t)
}

func TestService_GetHealth_AllServicesDown(t *testing.T) {
	// Arrange
	mockRepo := new(MockCheckerRepository)
	mockRepo.On("CheckDatabase").Return(errors.New("database connection failed"))
	mockRepo.On("CheckRedis").Return(errors.New("redis connection failed"))
	
	service := NewService(mockRepo)

	// Act
	result := service.GetHealth()

	// Assert
	assert.Equal(t, "degraded", result.Status)
	assert.Equal(t, "down", result.Services["database"])
	assert.Equal(t, "down", result.Services["redis"])
	mockRepo.AssertExpectations(t)
}

func TestNewService(t *testing.T) {
	// Arrange
	mockRepo := new(MockCheckerRepository)

	// Act
	service := NewService(mockRepo)

	// Assert
	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.repo)
}