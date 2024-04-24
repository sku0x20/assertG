package format

import (
	"fmt"
	"strings"
)

type Builder struct {
	sBuilder *strings.Builder
}

func (b *Builder) Value(value any) *Builder {
	b.sBuilder.WriteString(fmt.Sprintf("%v\n", value))
	return b
}

func (b *Builder) Message(message string) *Builder {
	b.sBuilder.WriteString(fmt.Sprintf("%s\n", message))
	return b
}

func (b *Builder) ToString() string {
	return b.sBuilder.String()
}

func Expected() *Builder {
	builder := &strings.Builder{}
	builder.WriteString("\nExpected\n")
	return &Builder{sBuilder: builder}
}
