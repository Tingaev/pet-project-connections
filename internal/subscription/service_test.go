package service

import (
    "testing"
    "context"

    m "github.com/tingaev/pet-project-connections/test/mocks"

    "github.com/google/uuid"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)


func TestCreate(t *testing.T) {
    // Arrange
    mockStorage := new(m.MockSubscriptionStorage)
    userUuid, _ := uuid.NewUUID()

    mockStorage.On("Create", mock.Anything, userUuid, userUuid).Return(nil).Once()
    service := NewService(mockStorage)

    // Act
    result := service.Create(context.Background(), userUuid, userUuid)

    // Assert
    require.Nil(t, result)
    mockStorage.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
    // Arrange
    mockStorage := new(m.MockSubscriptionStorage)
    userUuid, _ := uuid.NewUUID()

    mockStorage.On("Delete", mock.Anything, userUuid, userUuid).Return(nil).Once()
    service := NewService(mockStorage)

    // Act
    result := service.Delete(context.Background(), userUuid, userUuid)

    // Assert
    require.Nil(t, result)
    mockStorage.AssertExpectations(t)
}

func TestGetSubscribers(t *testing.T) {
    // Arrange
    mockStorage := new(m.MockSubscriptionStorage)
    userUuid, _ := uuid.NewUUID()
    userUuids := []uuid.UUID{userUuid, userUuid}
    offset := 0
    limit := 10

    mockStorage.On("GetSubscribers", mock.Anything, userUuid, offset, limit).Return(userUuids, nil).Once()
    service := NewService(mockStorage)

    // Act
    result, err := service.GetSubscribers(context.Background(), userUuid, offset, limit)

    // Assert
    require.Nil(t, err)
    assert.Equal(t, result, userUuids)
}

func TestGetSubscriptions(t *testing.T) {
    // Arrange
    mockStorage := new(m.MockSubscriptionStorage)
    userUuid, _ := uuid.NewUUID()
    userUuids := []uuid.UUID{userUuid, userUuid}
    offset := 0
    limit := 10

    mockStorage.On("GetSubscriptions", mock.Anything, userUuid, offset, limit).Return(userUuids, nil).Once()
    service := NewService(mockStorage)

    // Act
    result, err := service.GetSubscriptions(context.Background(), userUuid, offset, limit)

    // Assert
    require.Nil(t, err)
    assert.Equal(t, result, userUuids)
}

