package storage

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

