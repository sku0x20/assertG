package message

import (
	"fmt"
	"strings"
)

func NewMessage() *Message {
	return &Message{sb: &strings.Builder{}}
}

type Message struct {
	sb *strings.Builder
}

func (b *Message) Value(value any) *Message {
	return b.formattedStringWithNewLine("%v", value)
}

func (b *Message) Verb(message string) *Message {
	return b.formattedStringWithNewLine("%s", message)
}

func (b *Message) NewLine() *Message {
	return b.writeString("\n")
}

func (b *Message) ToString() string {
	return b.sb.String()
}

func (b *Message) formattedStringWithNewLine(format string, val any) *Message {
	return b.stringWithNewLine(fmt.Sprintf(format, val))
}

func (b *Message) stringWithNewLine(str string) *Message {
	b.writeString(str)
	return b.NewLine()
}

func (b *Message) writeString(str string) *Message {
	b.sb.WriteString(str)
	return b
}
