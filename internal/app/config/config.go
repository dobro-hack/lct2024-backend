package config

import (
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	DSN             string        `envconfig:"DSN" default:"root@tcp(localhost:3307)/lct2024?parseTime=true"`
	MaxIdleConns    int           `envconfig:"MAX_IDLE_CONNS" default:"10"`
	MaxOpenConns    int           `envconfig:"MAX_OPEN_CONNS" default:"10"`
	ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" default:"1m"`
}

var (
	cfg *Config
	mx  sync.RWMutex
)

func initConfig() (*Config, error) {
	godotenv.Load()
	cfg = &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Set main config
func Set(c Config) {
	mx.Lock()
	defer mx.Unlock()
	cfg = &c
}

// Get main config
func Get() Config {
	var err error
	mx.RLock()
	defer mx.RUnlock()
	if cfg == nil {
		cfg, err = initConfig()
		if err != nil {
			panic(err)
		}
	}
	return *cfg
}

// Reload main config
func Reload() Config {
	cfg, err := initConfig()
	if err != nil {
		return Get()
	}
	Set(*cfg)
	return Get()
}
