package commands

import (
	"context"
	"dddemo/internal/ddd"
	"dddemo/orders/internal/orders"
	"dddemo/orders/internal/payments"

	"github.com/pkg/errors"
)

type CreateOrderHandler struct {
	orderRepo       orders.Repository
	paymentRepo     payments.Repository
	domainPublisher ddd.EventPublisher
}

func NewCreateOrderHandler(orderRepo orders.Repository, paymentRepo payments.Repository, domainPublisher ddd.EventPublisher) CreateOrderHandler {
	return CreateOrderHandler{
		orderRepo:       orderRepo,
		paymentRepo:     paymentRepo,
		domainPublisher: domainPublisher,
	}
}

type CreateOrderSpec struct {
	ID         string
	CustomerID string
	PaymentID  string
	Items      []*orders.Item
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, spec CreateOrderSpec) error {
	order, err := orders.CreateOrder(spec.ID, spec.CustomerID, spec.PaymentID, spec.Items)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	if err = h.paymentRepo.Validate(ctx, order.PaymentID); err != nil {
		return errors.Wrap(err, "order payment confirmation")
	}

	if err = h.orderRepo.Save(ctx, *order); err != nil {
		return errors.Wrap(err, "order creation")
	}

	// call side-effect functions
	if err = h.domainPublisher.Publish(ctx, order.Events()...); err != nil {
		return err
	}

	return nil
}
