package exposure

import (
	"github.com/godome/godome/pkg/config"
	"github.com/godome/godome/pkg/logger"
)

type AdapterType string

type Adapter interface {
	GetType() AdapterType
	Config() config.Config
	Logger() logger.Logger
}
