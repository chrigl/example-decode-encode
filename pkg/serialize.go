package pkg

import (
	"strings"
)

type Foo struct {
	Name  string
	Space string

	NewField string
}

func Encode(f *Foo) string {
	var sb strings.Builder
	sb.WriteString(f.Name)
	sb.WriteString("_")
	sb.WriteString(f.Space)
	if f.NewField != "" {
		sb.WriteString("_")
		sb.WriteString(f.NewField)
	}
	return sb.String()
}

func Decode(data string) *Foo {
	tmp := strings.Split(data, "_")
	var newField string
	if len(tmp) == 3 {
		newField = tmp[2]
	}
	return &Foo{
		Name:     tmp[0],
		Space:    tmp[1],
		NewField: newField,
	}
}
