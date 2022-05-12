package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

const (
	ERR_ALREADY_EXISTS = "secret is already created"
)

type Config interface {
	Get(key string) string
	Set(key string, value string) error
}

type config struct {
	memCache map[string]string
}

var singletonConfig Config

func NewConfig() Config {
	if singletonConfig != nil {
		return singletonConfig
	}
	godotenv.Load()
	newConfig := &config{
		memCache: make(map[string]string),
	}
	singletonConfig = newConfig
	return newConfig
}

func (r *config) Get(key string) string {
	if r.memCache[key] != "" {
		return r.memCache[key]
	}
	return os.Getenv(key)
}

func (r *config) Set(key string, value string) error {
	if r.memCache[key] != "" {
		return errors.New(ERR_ALREADY_EXISTS)
	}
	r.memCache[key] = value
	return nil
}
