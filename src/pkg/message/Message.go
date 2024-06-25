package message

import (
	"assertG/src/pkg/message/verbs"
	"fmt"
	"strings"
)

type Message struct {
	sBuilder *strings.Builder
}

func (b *Message) Value(value any) *Message {
	return b.formattedStringWithNewLine("%v", value)
}

func (b *Message) Verb(message string) *Message {
	return b.formattedStringWithNewLine("%s", message)
}

func (b *Message) formattedStringWithNewLine(format string, val any) *Message {
	return b.stringWithNewLine(fmt.Sprintf(format, val))
}

func (b *Message) stringWithNewLine(str string) *Message {
	b.writeString(str)
	return b.newLine()
}

func (b *Message) newLine() *Message {
	return b.writeString("\n")
}

func (b *Message) writeString(str string) *Message {
	b.sBuilder.WriteString(str)
	return b
}

func (b *Message) ToString() string {
	return b.sBuilder.String()
}

func Expected() *Message {
	builder := newBuilder()
	builder.newLine()
	builder.Verb(verbs.Expected)
	return builder
}

func newBuilder() *Message {
	return &Message{sBuilder: &strings.Builder{}}
}
