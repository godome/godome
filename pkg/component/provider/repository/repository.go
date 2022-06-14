package repository

import (
	"github.com/godome/godome/pkg/component/adapter"
	"github.com/godome/godome/pkg/component/provider"
)

type Repository interface {
	provider.Provider
	GetAdapter(adapterName string) adapter.Adapter
	SetAdapter(adapter adapter.Adapter)
}

func NewRepository(name string) Repository {
	return &repository{
		Provider: provider.NewProvider(name),
		adapters: make(map[string]adapter.Adapter),
	}
}

type repository struct {
	provider.Provider
	adapters map[string]adapter.Adapter
}

func (r *repository) GetAdapter(adapterName string) adapter.Adapter {
	return r.adapters[adapterName]
}

func (r *repository) SetAdapter(adapter adapter.Adapter) {
	r.adapters[adapter.Metadata().GetName()] = adapter
}
