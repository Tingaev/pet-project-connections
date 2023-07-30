package mock

import (
    "context"
    "github.com/stretchr/testify/mock"
	"github.com/google/uuid"
)

type MockSubscriptionStorage struct {
    mock.Mock
}


func (m *MockSubscriptionStorage) Create(ctx context.Context, subscriberId uuid.UUID, targetId uuid.UUID) error {
    args := m.Called(ctx, subscriberId, targetId)
    return args.Error(0)
}

func (m *MockSubscriptionStorage) Delete(ctx context.Context, subscriberId uuid.UUID, targetId uuid.UUID) error {
    args := m.Called(ctx, subscriberId, targetId)	
	return args.Error(0)
}

func (m *MockSubscriptionStorage) GetSubscribers(ctx context.Context, id uuid.UUID, offset int, limit int) ([]uuid.UUID, error) {
    args := m.Called(ctx, id, offset, limit)
	return args.Get(0).([]uuid.UUID), args.Error(1)
}

func (m *MockSubscriptionStorage) GetSubscriptions(ctx context.Context, id uuid.UUID, offset int, limit int) ([]uuid.UUID, error) {
    args := m.Called(ctx, id, offset, limit)
	return args.Get(0).([]uuid.UUID), args.Error(1) 
}

