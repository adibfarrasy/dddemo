package payments

import (
	"context"
	"time"

	"github.com/rs/zerolog"
)

type Repository interface {
	Validate(ctx context.Context, paymentID string) error
}

type PaymentRepository struct {
	host   string
	port   string
	logger zerolog.Logger
}

var _ Repository = (*PaymentRepository)(nil)

func NewPaymentRepository(host string, port string, logger zerolog.Logger) PaymentRepository {
	return PaymentRepository{
		host:   host,
		port:   port,
		logger: logger,
	}
}

func (r PaymentRepository) Validate(ctx context.Context, paymentID string) error {
	// logic handwaving
	// this function makes an external call to payment domain
	// returns error if validation is unsuccessful
	r.logger.Info().Msg("Validating payment...")
	time.Sleep(1 * time.Second)
	r.logger.Info().Msg("Payment validated.")

	return nil
}
