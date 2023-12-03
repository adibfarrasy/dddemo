package invoices

import "time"

type Invoice struct {
	OrderID   string
	PaymentID string
	Amount    float64
	CreatedAt time.Time
}
