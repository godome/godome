package exposure

import (
	"github.com/godome/godome/pkg/config"
	"github.com/godome/godome/pkg/logger"
)

type ExposureType string

type Exposure interface {
	GetType() ExposureType
	Config() config.Config
	Logger() logger.Logger
}
