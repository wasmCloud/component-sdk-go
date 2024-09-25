package wasmcloud

import (
	"fmt"

	"go.wasmcloud.dev/component/gen/wasmcloud/secrets/reveal"
	"go.wasmcloud.dev/component/gen/wasmcloud/secrets/store"
)

// SecretGetAndReveal attempts to access a secret identified by the provided key
// and reveal it, returning the stored value as a slice of bytes.
func SecretGetAndReveal(key string) ([]byte, error) {
	res := store.Get(key)
	if res.IsErr() {
		return nil, fmt.Errorf("%v", res.Err())
	}

	opaqueSecret := *res.OK()
	defer opaqueSecret.ResourceDrop()

	revealed := reveal.Reveal(opaqueSecret)
	if s := revealed.String(); s != nil {
		return []byte(*s), nil
	}

	return revealed.Bytes().Slice(), nil
}
