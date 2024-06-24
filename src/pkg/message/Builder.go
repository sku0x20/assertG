package message

import (
	"assertG/src/pkg/message/messages"
	"fmt"
	"strings"
)

type Builder struct {
	sBuilder *strings.Builder
}

func (b *Builder) Value(value any) *Builder {
	b.sBuilder.WriteString(fmt.Sprintf("%v", value))
	b.newLine()
	return b
}

func (b *Builder) Message(message string) *Builder {
	b.sBuilder.WriteString(fmt.Sprintf("%s", message))
	b.newLine()
	return b
}

func (b *Builder) newLine() *Builder {
	b.sBuilder.WriteString("\n")
	return b
}

func (b *Builder) ToString() string {
	return b.sBuilder.String()
}

func Expected() *Builder {
	builder := newBuilder()
	builder.newLine()
	builder.Message(messages.Expected)
	return builder
}

func newBuilder() *Builder {
	return &Builder{sBuilder: &strings.Builder{}}
}
