package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env    string `yaml:"env" env-default:"local"`
	Server `yaml:"server"`
	Db     `yaml:"db"`
}

type Server struct {
	Address     string        `yaml:"address" env-default:":8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Db struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	Database string `yaml:"database" env-default:"postgres"`
	SSL      string `yaml:"ssl" env-default:"disable"`
}

func MustLoad() *Config {
	var viperCfg *viper.Viper
	var cfg Config

	viperCfg = viper.New()
	viperCfg.SetConfigName("config")
	viperCfg.AddConfigPath("./config")

	if err := viperCfg.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viperCfg.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
