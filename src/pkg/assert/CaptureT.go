package assert

import (
	"github.com/sku0x20/assertG/src/pkg/api"
)

func NewCaptureT(t api.T) *CaptureT {
	return &CaptureT{
		t,
	}
}

type CaptureT struct {
	t api.T
}

func (c *CaptureT) GetT() api.T {
	return c.t
}
