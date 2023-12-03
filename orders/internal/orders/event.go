package orders

type OrderCreated struct {
	Order *Order
}

func (OrderCreated) EventName() string { return "orders.OrderCreated" }
