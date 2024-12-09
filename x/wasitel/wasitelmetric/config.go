package wasitelmetric

import (
	"fmt"
	"net/url"
)

const (
	// DefaultPort is the default HTTP port of the collector.
	DefaultPort uint16 = 4318
	// DefaultHost is the host address the client will attempt
	// connect to if no collector address is provided.
	DefaultHost string = "localhost"
	// DefaultPath is a default URL path for endpoint that receives metrics.
	DefaultPath string = "/v1/metrics"
)

type config struct {
	Endpoint string
	Insecure bool
	Path     string
}

func newConfig(opts ...Option) config {
	cfg := config{
		Insecure: true,
		Endpoint: fmt.Sprintf("%s:%d", DefaultHost, DefaultPort),
		Path:     DefaultPath,
	}
	for _, opt := range opts {
		cfg = opt.apply(cfg)
	}
	return cfg
}

type Option interface {
	apply(config) config
}

func newWrappedOption(fn func(config) config) Option {
	return &wrappedOption{fn: fn}
}

type wrappedOption struct {
	fn func(config) config
}

func (o *wrappedOption) apply(cfg config) config {
	return o.fn(cfg)
}

func WithEndpoint(endpoint string) Option {
	return newWrappedOption(func(cfg config) config {
		cfg.Endpoint = endpoint
		return cfg
	})
}

func WithEndpointURL(eu string) Option {
	return newWrappedOption(func(cfg config) config {
		u, err := url.Parse(eu)
		if err != nil {
			return cfg
		}

		cfg.Endpoint = u.Host
		cfg.Path = u.Path
		if u.Scheme != "https" {
			cfg.Insecure = true
		}

		return cfg
	})
}
