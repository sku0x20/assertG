package message

import "assertG/src/pkg/message/verbs"

func Expected() *Message {
	builder := NewMessage()
	builder.NewLine()
	builder.Verb(verbs.Expected)
	return builder
}
