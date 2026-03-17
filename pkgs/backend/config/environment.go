package config

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	HOST string `mapstructure:"HOST"`
	PORT uint32 `mapstructure:"PORT"`

	NATS_PORT     string `mapstructure:"NATS_PORT"`
	NATS_HOST     string `mapstructure:"NATS_HOST"`
	EMBEDDED_NATS bool   `mapstructure:"EMBEDDED_NATS"`
}

func setDefaults() {
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", 7007)

	viper.SetDefault("NATS_PORT", 4222)
	viper.SetDefault("NATS_HOST", "nats")
	viper.SetDefault("EMBEDDED_NATS", true)
}

func GetEnvConfig() *EnvConfig {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	setDefaults()

	var envConfig EnvConfig
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}
	err = viper.Unmarshal(&envConfig)
	if err != nil {
		panic(err)
	}
	return &envConfig
}
