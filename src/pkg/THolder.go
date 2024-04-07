package pkg

import (
	"assertG/src/pkg/types"
)

var th = &tHolder{t: nil}

func RegisterT(t types.T) {
	th.t = t
}

func GetT() types.T {
	return th.t
}

func CleanT() {
	th.t = nil
}

type tHolder struct {
	t types.T
}
