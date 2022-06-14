package adapter

import (
	"github.com/godome/godome/pkg/component"
)

const componentType component.ComponentType = "adapter"

type Adapter interface {
	component.Component
}

type adapter struct {
	component.Component
}

func NewAdapter(name string) Adapter {
	return &adapter{
		Component: component.NewComponent(name, componentType),
	}
}
