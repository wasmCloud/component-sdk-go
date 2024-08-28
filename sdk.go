package component

// register wasm imports / exports
import (
	"embed"

	_ "github.com/lxfontes/component/gen/wasi/cli/environment"
)

//go:generate wit-bindgen-go generate --world imports --out gen ./wit

//go:embed wit/*
var Wit embed.FS

func Init() {
	// blank function. exercise the _ imports
}
