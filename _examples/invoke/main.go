//go:generate go run go.wasmcloud.dev/component/codegen --world example --out gen ./wit

package main

import (
	"github.com/wasmCloud/component-sdk-go/_examples/invoke/gen/example/invoker/invoker"
	"go.wasmcloud.dev/component/log/wasilog"
)

const InvokeResponse = "Hello from the invoker!"

func init() {
	invoker.Exports.Call = invokerCall
}

func invokerCall() string {
	logger := wasilog.ContextLogger("Call")
	logger.Info("Invoking function")
	return InvokeResponse
}

func main() {}
