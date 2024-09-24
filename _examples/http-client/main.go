package main

//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world example --out gen ./wit

import (
	"io"
	"net/http"

	"go.wasmcloud.dev/component/net/wasihttp"
)

var (
	wasiTransport = &wasihttp.Transport{}
	httpClient    = &http.Client{Transport: wasiTransport}
)

func init() {
	wasihttp.HandleFunc(proxyHandler)
}

func proxyHandler(w http.ResponseWriter, _ *http.Request) {
	url := "http://www.randomnumberapi.com/api/v1.0/random?min=100&max=1000&count=5"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "failed to create request", http.StatusBadGateway)
		return
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		http.Error(w, "failed to make outbound request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("X-Custom-Header", "proxied")
	w.WriteHeader(resp.StatusCode)

	_, _ = io.Copy(w, resp.Body)
}
func main() {}
