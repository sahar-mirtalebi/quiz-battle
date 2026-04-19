package config

import (
	"github.com/sahar-mirtalebi/quiz-battle/repository/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/service/authservice"
)

type HTTPConfig struct {
	Port int
}

type Config struct {
	Auth       authservice.Config
	HTTPConfig HTTPConfig
	Mysql      mysql.Config
}
