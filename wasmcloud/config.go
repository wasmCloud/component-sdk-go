package wasmcloud

import "go.wasmcloud.dev/component/gen/wasi/config/runtime"

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
