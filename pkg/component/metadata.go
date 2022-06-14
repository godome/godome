package component

type ComponentType string

type Metadata interface {
	GetName() string
	GetComponentType() ComponentType
}

func NewMetadata(name string, componentType ComponentType) Metadata {
	return &metadata{
		Name:          name,
		ComponentType: componentType,
	}
}

type metadata struct {
	Name          string
	ComponentType ComponentType
}

func (r *metadata) GetName() string {
	return r.Name
}

func (r *metadata) GetComponentType() ComponentType {
	return r.ComponentType
}
