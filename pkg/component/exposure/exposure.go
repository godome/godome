package exposure

import (
	"github.com/godome/godome/pkg/component"
)

const componentType component.ComponentType = "exposure"

type Exposure interface {
	component.Component
}

type exposure struct {
	component.Component
}

func NewExposure(name string) Exposure {
	return &exposure{
		Component: component.NewComponent(name, componentType),
	}
}
