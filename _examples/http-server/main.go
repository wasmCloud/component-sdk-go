//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world example --out gen ./wit

package main

import (
	"io"
	"net/http"
	"sort"
	"strings"

	"go.wasmcloud.dev/component/log/wasilog"
	"go.wasmcloud.dev/component/net/wasihttp"
)

const Index = `/error - return a 500 error
  /form - echo the fields of a POST request
  /headers - echo your user agent back as a server side header
  /post - echo the body of a POST request`

func init() {
	// We can't use http.ServeMux yet ( only symbol linking is supported in 'init' )
	wasihttp.HandleFunc(entryHandler)
}

func entryHandler(w http.ResponseWriter, r *http.Request) {
	logger := wasilog.ContextLogger("entryHandler")
	handlers := map[string]http.HandlerFunc{
		"/error":   errorHandler,
		"/form":    formHandler,
		"/headers": headersHandler,
		"/post":    postHandler,
	}

	logger.Info("Request received", "host", r.Host, "path", r.URL.Path, "agent", r.Header.Get("User-Agent"))

	if handler, ok := handlers[r.URL.Path]; ok {
		handler(w, r)
		return
	}

	var keys []string
	for k := range handlers {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	w.Header().Add("content-type", "text/plain")
	w.Header().Add("X-Requested-Path", r.URL.Path)
	w.Header().Add("X-Existing-Paths", strings.Join(keys, ","))
	_, _ = w.Write([]byte(Index))
}

func headersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Your-User-Agent", r.Header.Get("User-Agent"))

	_, _ = w.Write([]byte("Check headers!"))
}

func errorHandler(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "Something went wrong", http.StatusInternalServerError)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for key, values := range r.Form {
		_, _ = w.Write([]byte(key + ": " + strings.Join(values, ",") + "\n"))
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	logger := wasilog.ContextLogger("postHandler")

	w.WriteHeader(http.StatusOK)

	n, err := io.Copy(w, r.Body)
	if err != nil {
		logger.Error("Error copying body", "error", err)
		return
	}

	logger.Info("Copied body", "bytes", n)
}

func main() {}
