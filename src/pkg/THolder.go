package pkg

import "assertG/src/pkg/assert"

var th = &tHolder{t: nil}

func RegisterT(t assert.T) {
	th.t = t
}

func GetT() assert.T {
	return th.t
}

func CleanT() {
	th.t = nil
}

type tHolder struct {
	t assert.T
}
