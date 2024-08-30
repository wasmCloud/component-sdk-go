package lattice

import (
	"go.wasmcloud.dev/component/gen/wasmcloud/bus/lattice"
)

const DefaultLinkName = "default"

var (
	NewCallTargetInterface = lattice.NewCallTargetInterface
	SetLinkName            = lattice.SetLinkName
)
