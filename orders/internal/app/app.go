package app

import (
	"context"
	"dddemo/internal/ddd"
	"dddemo/orders/internal/app/commands"
	"dddemo/orders/internal/orders"
	"dddemo/orders/internal/payments"
)

type (
	App interface {
		CreateOrder(ctx context.Context, spec commands.CreateOrderSpec) error
	}

	Application struct {
		appCommands
		appQueries
	}

	appCommands struct {
		commands.CreateOrderHandler
	}

	appQueries struct {
		// not implemented
	}
)

var _ App = (*Application)(nil)

func New(o orders.Repository, p payments.Repository, d ddd.EventPublisher) *Application {
	return &Application{
		appCommands: appCommands{
			CreateOrderHandler: commands.NewCreateOrderHandler(o, p, d),
		},
		appQueries: appQueries{},
	}
}
