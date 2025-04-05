package assert_type

import "github.com/sku0x20/assertG/src/pkg/message"

type AssertType interface {
	FailWith(msg *message.Message)
}
