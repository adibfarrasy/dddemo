package config

import (
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
)

type (
	PGConfig struct {
		Conn string // `required:"true"`
	}

	RpcConfig struct {
		Host string `default:"0.0.0.0"`
		Port string `default:"8085"`
	}

	WebConfig struct {
		Host string `default:"0.0.0.0"`
		Port string `default:"8080"`
	}

	AppConfig struct {
		Environment     string
		LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		PG              PGConfig
		Rpc             RpcConfig
		Web             WebConfig
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
