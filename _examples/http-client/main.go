package main

//go:generate wit-bindgen-go generate --world example --out gen ./wit

import (
	"io"
	"net/http"

	"go.wasmcloud.dev/component/log/wasilog"
	"go.wasmcloud.dev/component/net/wasihttp"
)

var (
	wasiTransport = &wasihttp.Transport{}
	httpClient    = &http.Client{Transport: wasiTransport}
)

func init() {
	wasihttp.HandleFunc(proxyHandler)
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	logger := wasilog.ContextLogger("proxyHandler")

	logger.Info("creating request")
	url := "http://www.randomnumberapi.com/api/v1.0/random?min=100&max=1000&count=5"
	url = "http://localhost:8081/"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "failed to create request", http.StatusBadGateway)
		return
	}

	logger.Info("executing request")
	resp, err := httpClient.Do(req)
	if err != nil {
		http.Error(w, "failed to make outbound request", http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	logger.Info("proxying")
	w.Header().Set("X-Custom-Header", "proxied")
	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}
func main() {}
