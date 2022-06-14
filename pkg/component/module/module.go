package module

import (
	"github.com/godome/godome/pkg/component"
	"github.com/godome/godome/pkg/component/provider"
)

const componentType component.ComponentType = "module"

type Module interface {
	component.Component
	SetProvider(provider provider.Provider)
	GetProvider(name string) provider.Provider
}

func NewModule(name string) Module {
	return &module{
		Component: component.NewComponent(name, componentType),
		providers: make(map[string]provider.Provider),
	}
}

type module struct {
	component.Component
	providers map[string]provider.Provider
}

func (r *module) GetProvider(name string) provider.Provider {
	return r.providers[name]
}

func (r *module) SetProvider(provider provider.Provider) {
	r.providers[provider.Metadata().GetName()] = provider
}
