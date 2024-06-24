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
	return b.formattedStringWithNewLine("%v", value)
}

func (b *Builder) Message(message string) *Builder {
	return b.formattedStringWithNewLine("%s", message)
}

func (b *Builder) formattedStringWithNewLine(format string, val any) *Builder {
	return b.stringWithNewLine(fmt.Sprintf(format, val))
}

func (b *Builder) stringWithNewLine(str string) *Builder {
	b.writeString(str)
	return b.newLine()
}

func (b *Builder) newLine() *Builder {
	return b.writeString("\n")
}

func (b *Builder) writeString(str string) *Builder {
	b.sBuilder.WriteString(str)
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
