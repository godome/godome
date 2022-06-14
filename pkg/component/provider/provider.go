package provider

import (
	"github.com/godome/godome/pkg/component"
)

const componentType component.ComponentType = "provider"

type Provider interface {
	component.Component
}

func NewProvider(name string) Provider {
	return &provider{
		Component: component.NewComponent(name, componentType),
	}
}

type provider struct {
	component.Component
}
