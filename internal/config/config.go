// Package config loads and represents application configuration.
package config

import (
	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/redis"
	"github.com/sahar-mirtalebi/quiz-battle/internal/auth/token"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"
)

type HTTPConfig struct {
	Port int `koanf:"port"`
}

type Config struct {
	JWT             authservice.Config     `koanf:"jwt"`
	HTTPConfig      HTTPConfig             `koanf:"http_config"`
	Mysql           mysql.Config           `koanf:"mysql"`
	Redis           redis.Config           `koanf:"redis"`
	MatchingService matchingservice.Config `koanf:"matching_service"`
}
