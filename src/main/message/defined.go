package message

import "github.com/sku0x20/assertG/src/main/message/verbs"

func Expected() *Message {
	return NewMessage().
		NewLine().
		Verb(verbs.Expected)
}
