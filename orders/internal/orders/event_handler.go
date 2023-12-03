package orders

import (
	"context"
	"dddemo/internal/ddd"
)

type DomainEventHandler interface {
	OnOrderCreated(ctx context.Context, event ddd.Event) error
}

type IgnoreUnimplementedDomainEvents struct{}

var _ DomainEventHandler = (*IgnoreUnimplementedDomainEvents)(nil)

func (IgnoreUnimplementedDomainEvents) OnOrderCreated(ctx context.Context, event ddd.Event) error {
	return nil
}
