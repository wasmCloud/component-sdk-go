package wasitelmetric

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go.wasmcloud.dev/component/net/wasihttp"
	"go.wasmcloud.dev/component/x/wasitel/wasitelmetric/internal/types"
)

type clienti interface {
	UploadMetrics(context.Context, *types.ResourceMetrics) error
	Shutdown(context.Context) error
}

func newClient(opts ...Option) *client {
	cfg := newConfig(opts...)

	wasiTransport := &wasihttp.Transport{}
	httpClient := &http.Client{Transport: wasiTransport}

	return &client{
		config:     cfg,
		httpClient: httpClient,
	}
}

type client struct {
	config     config
	httpClient *http.Client
}

func (c *client) UploadMetrics(ctx context.Context, rm *types.ResourceMetrics) error {
	export := &types.ExportMetricsServiceRequest{
		ResourceMetrics: []*types.ResourceMetrics{rm},
	}

	body, err := json.Marshal(export)
	if err != nil {
		return fmt.Errorf("failed to serialize export request to JSON: %w", err)
	}

	u := c.getUrl()
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request to %q: %w", u.String(), err)
	}
	req.Header.Set("Content-Type", "application/json")

	_, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to request %q: %w", u.String(), err)
	}

	return nil
}

// Shutdown shuts down the client, freeing all resources.
func (c *client) Shutdown(ctx context.Context) error {
	c.httpClient = nil
	return ctx.Err()
}

func (c *client) getUrl() url.URL {
	scheme := "http"
	if !c.config.Insecure {
		scheme = "https"
	}
	u := url.URL{
		Scheme: scheme,
		Host:   c.config.Endpoint,
		Path:   c.config.Path,
	}
	return u
}
