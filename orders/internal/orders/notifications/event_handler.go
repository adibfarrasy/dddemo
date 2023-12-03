package notifications

import (
	"context"
	"dddemo/internal/ddd"
	"dddemo/orders/internal/orders"
)

type NotificationHandler struct {
	repo Repository
	orders.IgnoreUnimplementedDomainEvents
}

func NewNotificationHandler(repo Repository) NotificationHandler {
	return NotificationHandler{
		repo: repo,
	}
}

var _ orders.DomainEventHandler = (*NotificationHandler)(nil)

func (h NotificationHandler) OnOrderCreated(ctx context.Context, event ddd.Event) error {
	ev := event.(*orders.OrderCreated)
	notification := NotifPayload{
		Title: ev.Order.ID,
	}
	h.repo.Send(ctx, notification)

	return nil
}
