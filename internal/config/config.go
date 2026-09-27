// Package config loads and represents application configuration.
package config

import (
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/redis"
	authservice "github.com/sahar-mirtalebi/quiz-battle/internal/auth/token"
	matchingrepo "github.com/sahar-mirtalebi/quiz-battle/internal/matching/repository/redis"
	matchingservice "github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"
	presenceservice "github.com/sahar-mirtalebi/quiz-battle/internal/presence/service"
)

type Application struct {
	GracefullShutdownTimeout time.Duration `koanf:"gracefull_shutdown_timeout"`
}

type Scheduler struct {
	MatchPlayersInterval time.Duration `koanf:"match_players_interval"`
	MatchPlayersTimeout  time.Duration `koanf:"match_players_timeout"`
}

type HTTPConfig struct {
	Port int `koanf:"port"`
}

type Config struct {
	MatchingRepository matchingrepo.Config    `koanf:"matching_repository"`
	JWT                authservice.Config     `koanf:"jwt"`
	HTTPConfig         HTTPConfig             `koanf:"http_config"`
	Mysql              mysql.Config           `koanf:"mysql"`
	Redis              redis.Config           `koanf:"redis"`
	MatchingService    matchingservice.Config `koanf:"matching_service"`
	Application        Application            `koanf:"application"`
	Scheduler          Scheduler              `koanf:"scheduler"`
	PresenceService    presenceservice.Config `koanf:"presence_service"`
}
