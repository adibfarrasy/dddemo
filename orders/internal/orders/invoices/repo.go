package invoices

import (
	"context"
	"database/sql"
	"time"

	"github.com/rs/zerolog"
)

type Repository interface {
	Save(ctx context.Context, invoice Invoice) error
}

type InvoiceRepository struct {
	db     *sql.DB
	logger zerolog.Logger
}

var _ Repository = (*InvoiceRepository)(nil)

func NewInvoiceRepository(db *sql.DB, logger zerolog.Logger) InvoiceRepository {
	return InvoiceRepository{
		db:     db,
		logger: logger,
	}
}

func (r InvoiceRepository) Save(ctx context.Context, invoice Invoice) error {
	// logic handwaving
	// this function inserts a new invoice upon order creation
	// to the invoices table in orders db
	r.logger.Info().Msg("Saving invoice...")
	time.Sleep(1 * time.Second)
	r.logger.Info().Msg("Invoice saved.")

	return nil
}
