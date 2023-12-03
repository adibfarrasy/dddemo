package orders

import (
	"context"
	"dddemo/internal/ddd"
	"dddemo/internal/module"
	"dddemo/orders/internal/app"
	"dddemo/orders/internal/app/commands"
	"dddemo/orders/internal/orders"
	"dddemo/orders/internal/orders/invoices"
	"dddemo/orders/internal/orders/notifications"
	"dddemo/orders/internal/payments"
	"fmt"
	"net/http"
)

type Module struct{}

func (m Module) Setup(ctx context.Context, mono module.Monolith) error {
	logger := mono.GetLogger()

	// setup "driven" adapters
	domainDispatcher := ddd.NewEventDispatcher()
	orderRepo := orders.NewOrderRepository(nil, logger)
	invoiceRepo := invoices.NewInvoiceRepository(nil, logger)
	notifRepo := notifications.NewNotificationRepository("channel_id", logger)
	paymentExtRepo := payments.NewPaymentRepository("", "", logger)

	// setup application
	application := app.New(orderRepo, paymentExtRepo, domainDispatcher)

	// setup "driver" adapters
	notifEventHandler := notifications.NewNotificationHandler(notifRepo)
	invoiceEventHandler := invoices.NewInvoiceHandler(invoiceRepo)
	subscribeOrderCreatedEvents(domainDispatcher, invoiceEventHandler, notifEventHandler)

	mono.GetMux().Post("/orders", func(w http.ResponseWriter, r *http.Request) {
		if err := application.CreateOrder(ctx, commands.CreateOrderSpec{
			// payload hardcoded for demo purposes
			ID:         "id",
			PaymentID:  "payment_id",
			CustomerID: "customer_id",
			Items: []*orders.Item{
				{ProductID: "product_id"},
			},
		}); err != nil {
			logger.Info().Msg(fmt.Sprintf("error, %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte("An order is created."))
		return
	})

	return nil
}

func subscribeOrderCreatedEvents(dispathcer *ddd.EventDispatcher, invoice orders.DomainEventHandler, notification orders.DomainEventHandler) {
	dispathcer.Subscribe(orders.OrderCreated{}, invoice.OnOrderCreated)
	dispathcer.Subscribe(orders.OrderCreated{}, notification.OnOrderCreated)
}
