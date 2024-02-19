package service

import (
	"context"

	"github.com/google/uuid"
)

type Subscription interface {
	Create(ctx context.Context, subscriberId uuid.UUID, targetId uuid.UUID) error
	Delete(ctx context.Context, subscriberId uuid.UUID, targetId uuid.UUID) error

	GetSubscribers(ctx context.Context, id uuid.UUID, offset int, limit int) ([]uuid.UUID, error)
	GetSubscriptions(ctx context.Context, id uuid.UUID, offset int, limit int) ([]uuid.UUID, error)
}

type service struct {
	subscriptionStorage Subscription
}

func NewService(s Subscription) service {
	service := service{
		subscriptionStorage: s,
	}
	return service
}

func (s *service) Create(ctx context.Context, subscriberId uuid.UUID, targetId uuid.UUID) error {
	return s.subscriptionStorage.Create(ctx, subscriberId, targetId)
}

func (s *service) Delete(ctx context.Context, subscriberId uuid.UUID, targetId uuid.UUID) error {
	return s.subscriptionStorage.Delete(ctx, subscriberId, targetId)
}

func (s *service) GetSubscribers(ctx context.Context, id uuid.UUID, offset int, limit int) ([]uuid.UUID, error) {
	return s.subscriptionStorage.GetSubscribers(ctx, id, offset, limit)
}

func (s *service) GetSubscriptions(ctx context.Context, id uuid.UUID, offset int, limit int) ([]uuid.UUID, error) {
	return s.subscriptionStorage.GetSubscriptions(ctx, id, offset, limit)
}
