package assertions

import "github.com/sku0x20/assertG/src/pkg/message"

type Assertion interface {
	FailWith(msg *message.Message)
}
