package wasiteltrace

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go.wasmcloud.dev/component/net/wasihttp"
	"go.wasmcloud.dev/component/x/wasitel/wasiteltrace/internal/types"
)

type client struct {
	config     config
	httpClient *http.Client
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

func (c *client) UploadTraces(ctx context.Context, spans []*types.ResourceSpans) error {
	if len(spans) == 0 {
		return nil
	}

	export := &types.ExportTraceServiceRequest{
		ResourceSpans: spans,
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
