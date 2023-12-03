package orders

import "dddemo/internal/ddd"

func CreateOrder(id, customerID, paymentID string, items []*Item) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	if paymentID == "" {
		return nil, ErrPaymentIDCannotBeBlank
	}

	order := &Order{
		AggregateBase: ddd.AggregateBase{
			ID: id,
		},
		CustomerID: customerID,
		PaymentID:  paymentID,
		Items:      items,
		Status:     OrderIsPending,
	}

	order.AddEvent(&OrderCreated{
		Order: order,
	})

	return order, nil
}
