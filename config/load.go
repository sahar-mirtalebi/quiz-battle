package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"

	"github.com/knadh/koanf/providers/env/v2"
)

func Load(configPath string) *Config {
	var k = koanf.New(".")

	_ = godotenv.Load()

	k.Load(confmap.Provider(defaultConfig, "."), nil)

	b, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	expanded := os.ExpandEnv(string(b))

	k.Load(rawbytes.Provider([]byte(expanded)), yaml.Parser())

	k.Load(env.Provider(".", env.Opt{
		Prefix: "GAMEAPP_",
		TransformFunc: func(k, v string) (string, any) {
			k = strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(k, "GAMEAPP_")), "_", ".")

			return k, v
		},
	}), nil)

	var cfg Config

	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
