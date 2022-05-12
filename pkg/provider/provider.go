package provider

type ProviderType string

type Provider interface {
	GetType() ProviderType
}
