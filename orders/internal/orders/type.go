package orders

import (
	"dddemo/internal/ddd"
)

type OrderStatus string

const (
	OrderIsPending   OrderStatus = "pending"
	OrderIsInProcess OrderStatus = "in-progress"
	OrderIsCompleted OrderStatus = "completed"
	OrderIsCancelled OrderStatus = "cancelled"
)

type Order struct {
	ddd.AggregateBase
	CustomerID string
	PaymentID  string
	InvoiceID  string
	ShoppingID string
	Items      []*Item
	Status     OrderStatus
	Amount     float64
}

type Item struct {
	ProductID   string
	StoreID     string
	StoreName   string
	ProductName string
	Price       float64
	Quantity    int
}
