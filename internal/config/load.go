package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func Load(configPath string) *Config {
	var k = koanf.New(".")

	_ = godotenv.Load()

	if err := k.Load(confmap.Provider(defaultConfig, "."), nil); err != nil {
		panic(richerror.New().WithError(err).WithKind(richerror.KindUnexpected))
	}

	b, err := os.ReadFile(configPath)
	if err != nil {
		panic(richerror.New().WithError(err).WithKind(richerror.KindUnexpected))
	}

	expanded := os.ExpandEnv(string(b))

	if err := k.Load(rawbytes.Provider([]byte(expanded)), yaml.Parser()); err != nil {
		panic(richerror.New().WithError(err).WithKind(richerror.KindUnexpected))
	}

	if err := k.Load(env.Provider(".", env.Opt{
		Prefix: "GAMEAPP_",
		TransformFunc: func(k, v string) (string, any) {
			k = strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(k, "GAMEAPP_")), "_", ".")

			return k, v
		},
	}), nil); err != nil {
		panic(richerror.New().WithError(err).WithKind(richerror.KindUnexpected))
	}

	var cfg Config

	if err := k.Unmarshal("", &cfg); err != nil {
		panic(richerror.New().WithError(err).WithKind(richerror.KindUnexpected))
	}

	return &cfg
}
