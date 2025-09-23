package health

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDatabase is a mock implementation of configs.DatabaseInterface
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockDatabase) Close() error {
	args := m.Called()
	return args.Error(0)
}

// MockRedis is a mock implementation of configs.RedisInterface
type MockRedis struct {
	mock.Mock
}

func (m *MockRedis) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRedis) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestCheckerRepository_CheckDatabase_Success(t *testing.T) {
	// Arrange
	mockDB := new(MockDatabase)
	mockRedis := new(MockRedis)
	mockDB.On("Ping").Return(nil)

	repo := NewCheckerRepository(mockDB, mockRedis)

	// Act
	err := repo.CheckDatabase()

	// Assert
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestCheckerRepository_CheckDatabase_Failure(t *testing.T) {
	// Arrange
	mockDB := new(MockDatabase)
	mockRedis := new(MockRedis)
	mockDB.On("Ping").Return(errors.New("database connection failed"))

	repo := NewCheckerRepository(mockDB, mockRedis)

	// Act
	err := repo.CheckDatabase()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database connection failed")
	mockDB.AssertExpectations(t)
}

func TestCheckerRepository_CheckRedis_Success(t *testing.T) {
	// Arrange
	mockDB := new(MockDatabase)
	mockRedis := new(MockRedis)
	mockRedis.On("Ping").Return(nil)

	repo := NewCheckerRepository(mockDB, mockRedis)

	// Act
	err := repo.CheckRedis()

	// Assert
	assert.NoError(t, err)
	mockRedis.AssertExpectations(t)
}

func TestCheckerRepository_CheckRedis_Failure(t *testing.T) {
	// Arrange
	mockDB := new(MockDatabase)
	mockRedis := new(MockRedis)
	mockRedis.On("Ping").Return(errors.New("redis connection failed"))

	repo := NewCheckerRepository(mockDB, mockRedis)

	// Act
	err := repo.CheckRedis()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "redis connection failed")
	mockRedis.AssertExpectations(t)
}

func TestNewCheckerRepository(t *testing.T) {
	// Arrange
	mockDB := new(MockDatabase)
	mockRedis := new(MockRedis)

	// Act
	repo := NewCheckerRepository(mockDB, mockRedis)

	// Assert
	assert.NotNil(t, repo)
	assert.IsType(t, &CheckerRepository{}, repo)
}