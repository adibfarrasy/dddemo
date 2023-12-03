package notifications

import (
	"context"
	"time"

	"github.com/rs/zerolog"
)

type Repository interface {
	Send(ctx context.Context, notification NotifPayload) error
}

type NotificationRepository struct {
	channelID string
	logger    zerolog.Logger
}

var _ Repository = (*NotificationRepository)(nil)

func NewNotificationRepository(channelID string, logger zerolog.Logger) NotificationRepository {
	return NotificationRepository{
		channelID: channelID,
		logger:    logger,
	}
}

func (r NotificationRepository) Send(ctx context.Context, notification NotifPayload) error {
	// logic handwaving
	// the handler sends a notification that an order is created
	// via external communication channel e.g. SMS
	r.logger.Info().Msg("Sending order created notification...")
	time.Sleep(1 * time.Second)
	r.logger.Info().Msg("Notification sent.")

	return nil
}
