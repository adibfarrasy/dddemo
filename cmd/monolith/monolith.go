package monolith

import (
	"context"
	"database/sql"
	"dddemo/internal/config"
	"dddemo/internal/module"
	"dddemo/internal/waiter"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type App struct {
	Cfg     config.AppConfig
	DB      *sql.DB
	Logger  zerolog.Logger
	Mux     *chi.Mux
	RPC     *grpc.Server
	Modules []module.Module
	Waiter  waiter.Waiter
}

var _ module.Monolith = (*App)(nil)

func (a *App) SetupModules() error {
	for _, m := range a.Modules {
		if err := m.Setup(a.Waiter.Context(), a); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) WaitForWeb(ctx context.Context) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", a.Cfg.Web.Host, a.Cfg.Web.Port),
		Handler: a.Mux,
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		fmt.Println("web server started")
		defer fmt.Println("web server shutdown")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	g.Go(func() error {
		<-gCtx.Done()
		fmt.Println("shutting down web server...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	return g.Wait()
}

func (a *App) GetMux() *chi.Mux {
	return a.Mux
}

func (a *App) GetLogger() zerolog.Logger {
	return a.Logger
}
