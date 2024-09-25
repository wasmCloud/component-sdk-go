package wasmcloud

import "go.wasmcloud.dev/component/gen/wasi/config/runtime"

// GetConfigOrDefault tries to get a configuration value by provided key using
// [wasi:config/store.get] and falling back to the provided defaultValue if a
// configuration value by provided key is not found.
//
// [wasi:config/store.get]: https://github.com/WebAssembly/wasi-runtime-config/blob/f4d699bc6dd77adad99fa1a2246d482225ec6485/wit/store.wit#L17-L24
func GetConfigOrDefault(key string, defaultValue string) string {
	res := runtime.Get(key)
	if res.IsOK() {
		opt := *res.OK()
		if !opt.None() {
			return *opt.Some()
		}
	}

	return defaultValue
}
