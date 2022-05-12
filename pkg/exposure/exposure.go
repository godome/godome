package exposure

type ExposureType string

type Exposure interface {
	GetType() ExposureType
}
