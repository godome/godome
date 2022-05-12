package module

import (
	"github.com/godome/godome/pkg/config"
	"github.com/godome/godome/pkg/provider"
)

type Module interface {
	GetName() string
	Config() config.Config
	SetProvider(provider provider.Provider)
	GetProvider(providerType provider.ProviderType) provider.Provider
}

type module struct {
	name      string
	config    config.Config
	providers map[provider.ProviderType]provider.Provider
}

func NewModule(name string) Module {
	return &module{
		name:      name,
		config:    config.NewConfig(),
		providers: make(map[provider.ProviderType]provider.Provider),
	}
}

func (r *module) GetName() string {
	return r.name
}

func (r *module) Config() config.Config {
	return r.config
}

func (r *module) GetProvider(providerType provider.ProviderType) provider.Provider {
	return r.providers[providerType]
}

func (r *module) SetProvider(provider provider.Provider) {
	r.providers[provider.GetType()] = provider
}
