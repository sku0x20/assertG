package assertions

import "assertG/src/pkg/message"

type Assertion interface {
	FailWith(msg *message.Message)
}
