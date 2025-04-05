package assert_type

import (
	"github.com/sku0x20/assertG/src/main/message"
	"testing"
)

func NewSoftAssert(t *testing.T) *SoftAssert {
	return &SoftAssert{t, false}
}

type SoftAssert struct {
	t      *testing.T
	failed bool
}

func (s *SoftAssert) FailWith(msg *message.Message) {
	s.failed = true
	s.t.Helper()
	s.t.Logf(msg.ToString())
}

func (s *SoftAssert) Failed() bool {
	return s.failed
}
