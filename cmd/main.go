package main

import (
	"dddemo/cmd/monolith"
	"dddemo/internal/config"
	"dddemo/internal/module"
	"dddemo/internal/waiter"
	"dddemo/orders"
	"fmt"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	// init config
	var cfg config.AppConfig
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	// init monolith
	mono := monolith.App{Cfg: cfg}

	// setup DB
	// mono.DB, err = sql.Open("pgx", cfg.PG.Conn)
	// if err != nil {
	// 	return err
	// }
	// defer func(db *sql.DB) {
	// 	err := db.Close()
	// 	if err != nil {
	// 		return
	// 	}
	// }(mono.DB)

	// setup logger
	mono.Logger = initLogger(cfg.Environment)

	// setup gRPC server
	server := grpc.NewServer()
	reflection.Register(server)
	mono.RPC = server

	// setup http server
	mono.Mux = chi.NewMux()

	mono.Waiter = waiter.New()

	// load modules
	mono.Modules = []module.Module{
		orders.Module{},
	}

	if err = mono.SetupModules(); err != nil {
		return err
	}

	fmt.Println("started application")
	defer fmt.Println("stopped application")

	mono.Waiter.Add(
		mono.WaitForWeb,
	)

	return mono.Waiter.Wait()
}

func initLogger(env string) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	switch env {
	case "production":
		return zerolog.New(os.Stdout).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Logger()
	default:
		return zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = "03:04:05.000PM"
		})).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Logger()
	}
}
