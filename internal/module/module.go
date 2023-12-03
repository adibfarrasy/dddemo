package module

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type Monolith interface {
	SetupModules() error
	WaitForWeb(ctx context.Context) error
	GetMux() *chi.Mux
	GetLogger() zerolog.Logger
}

type Module interface {
	Setup(ctx context.Context, mono Monolith) error
}
