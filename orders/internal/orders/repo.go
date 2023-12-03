package orders

import (
	"context"
	"database/sql"
	"time"

	"github.com/rs/zerolog"
)

type Repository interface {
	Save(ctx context.Context, order Order) error
}

type OrderRepository struct {
	db     *sql.DB
	logger zerolog.Logger
}

var _ Repository = (*OrderRepository)(nil)

func NewOrderRepository(db *sql.DB, logger zerolog.Logger) OrderRepository {
	return OrderRepository{
		db:     db,
		logger: logger,
	}
}

func (r OrderRepository) Save(ctx context.Context, order Order) error {
	// logic handwaving
	// this function inserts a new order to the database
	r.logger.Info().Msg("Saving order...")
	time.Sleep(1 * time.Second)
	r.logger.Info().Msg("Order saved.")

	return nil
}
