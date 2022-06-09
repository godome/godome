package repository

import (
	"github.com/godome/godome/pkg/adapter"
	"github.com/godome/godome/pkg/config"
	"github.com/godome/godome/pkg/logger"
)

type Repository interface {
	Config() config.Config
	Logger() logger.Logger
	GetAdapter(adapterType adapter.AdapterType) adapter.Adapter
	SetAdapter(adapter adapter.Adapter)
}

type repository struct {
	config   config.Config
	adapters map[adapter.AdapterType]adapter.Adapter
}

func NewRepository() Repository {
	return &repository{
		config:   config.NewConfig(),
		adapters: make(map[adapter.AdapterType]adapter.Adapter),
	}
}

func (r *repository) Config() config.Config {
	return r.config
}

func (r *repository) Logger() logger.Logger {
	return logger.GetLogger()
}

func (r *repository) GetAdapter(adapterType adapter.AdapterType) adapter.Adapter {
	return r.adapters[adapterType]
}

func (r *repository) SetAdapter(adapter adapter.Adapter) {
	r.adapters[adapter.GetType()] = adapter
}
