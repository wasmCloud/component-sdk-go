package wasitime

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	wasiNow := Now()
	runtimeNow := time.Now()
	t.Log("wasiNow: ", wasiNow)
	t.Log("runtimeNow: ", runtimeNow)
}
