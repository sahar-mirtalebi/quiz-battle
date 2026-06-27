package config

import (
	"github.com/sahar-mirtalebi/quiz-battle/repository/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/service/authservice"
)

type HTTPConfig struct {
	Port int `koanf:"port"`
}

type Config struct {
	JWT        authservice.Config `koanf:"jwt"`
	HTTPConfig HTTPConfig         `koanf:"http_config"`
	Mysql      mysql.Config       `koanf:"mysql"`
}
