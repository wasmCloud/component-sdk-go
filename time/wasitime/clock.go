package wasitime

import (
	"time"

	wallclock "go.wasmcloud.dev/component/gen/wasi/clocks/wall-clock"
)

func Now() time.Time {
	res := wallclock.Now()
	return time.Unix(int64(res.Seconds), int64(res.Nanoseconds))
}
