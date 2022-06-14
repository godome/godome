package component

import (
	"github.com/godome/godome/pkg/config"
	"github.com/godome/godome/pkg/logger"
)

type Component interface {
	Config() config.Config
	Logger() logger.Logger
	Metadata() Metadata
}

type component struct {
	metadata Metadata
	config   config.Config
	logger   logger.Logger
}

func NewComponent(name string, componentType ComponentType) Component {
	return &component{
		metadata: NewMetadata(
			name,
			componentType,
		),
		logger: logger.GetLogger(),
		config: config.NewConfig(),
	}
}

func (r *component) Logger() logger.Logger {
	return r.logger
}

func (r *component) Config() config.Config {
	return r.config
}

func (r *component) Metadata() Metadata {
	return r.metadata
}
