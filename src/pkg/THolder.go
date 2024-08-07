package pkg

import (
	"github.com/sku0x20/assertG/src/pkg/types"
)

var th = &tHolder{t: nil}

func RegisterT(t types.T) {
	th.t = t
}

func GetT() types.T {
	if th.t == nil {
		panic("t is not set")
	}
	return th.t
}

func CleanT() {
	th.t = nil
}

func EnableAsserts(t types.T) func() {
	RegisterT(t)
	return func() {
		CleanT()
	}
}

type tHolder struct {
	t types.T
}
