package pkg

import (
	"github.com/sku0x20/assertG/src/pkg/api"
)

var th = &tHolder{t: nil}

func RegisterT(t api.T) {
	th.t = t
}

func GetT() api.T {
	if th.t == nil {
		panic("t is not set")
	}
	return th.t
}

func CleanT() {
	th.t = nil
}

func EnableAsserts(t api.T) func() {
	RegisterT(t)
	return func() {
		CleanT()
	}
}

type tHolder struct {
	t api.T
}
