package lattice

import (
	"go.wasmcloud.dev/component/gen/wasmcloud/bus/lattice"
)

const DefaultLinkName = "default"

type CallTargetInterface = lattice.CallTargetInterface

var (
	NewCallTargetInterface = lattice.NewCallTargetInterface
	SetLinkName            = lattice.SetLinkName
)
