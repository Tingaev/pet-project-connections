package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/tingaev/pet-project-connections/internal/storage"
)


type Subscriptions interface {
	storage.Subscription // TODO: когда появятся отличия в их интерфейсе, тогда нужно будет это убрать
}


type service struct {
	subscriptionStorage storage.Subscription
}


func NewService(s storage.Subscription) service {
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
