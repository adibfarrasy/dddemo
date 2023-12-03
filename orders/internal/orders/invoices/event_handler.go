package invoices

import (
	"context"
	"dddemo/internal/ddd"
	"dddemo/orders/internal/orders"
	"time"
)

type InvoiceHandler struct {
	repo Repository
	orders.IgnoreUnimplementedDomainEvents
}

func NewInvoiceHandler(repo Repository) InvoiceHandler {
	return InvoiceHandler{
		repo: repo,
	}
}

var _ orders.DomainEventHandler = (*InvoiceHandler)(nil)

func (h InvoiceHandler) OnOrderCreated(ctx context.Context, event ddd.Event) error {
	ev := event.(*orders.OrderCreated)
	invoice := Invoice{
		OrderID:   ev.Order.ID,
		PaymentID: ev.Order.PaymentID,
		Amount:    ev.Order.Amount,
		CreatedAt: time.Now(),
	}
	h.repo.Save(ctx, invoice)

	return nil
}
