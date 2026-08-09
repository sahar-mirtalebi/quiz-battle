// Package config loads and represents application configuration.
package config

import (
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/redis"
	authservice "github.com/sahar-mirtalebi/quiz-battle/internal/auth/token"
	matchingservice "github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"
)

type Application struct {
	GracefullShutdownTimeout time.Duration `koanf:"gracefull_shutdown_timeout"`
}

type Scheduler struct {
	MatchPlayersInterval time.Duration `koanf:"match_players_interval"`
}

type HTTPConfig struct {
	Port int `koanf:"port"`
}

type Config struct {
	JWT             authservice.Config     `koanf:"jwt"`
	HTTPConfig      HTTPConfig             `koanf:"http_config"`
	Mysql           mysql.Config           `koanf:"mysql"`
	Redis           redis.Config           `koanf:"redis"`
	MatchingService matchingservice.Config `koanf:"matching_service"`
	Application     Application            `koanf:"application"`
	Scheduler       Scheduler              `koanf:"scheduler"`
}
