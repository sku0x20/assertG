package message

import "assertG/src/pkg/message/verbs"

func Expected() *Message {
	return NewMessage().
		NewLine().
		Verb(verbs.Expected)
}
