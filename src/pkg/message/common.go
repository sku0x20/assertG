package message

import "assertG/src/pkg/message/verbs"

func Expected() *Message {
	message := NewMessage()
	message.NewLine()
	message.Verb(verbs.Expected)
	return message
}
