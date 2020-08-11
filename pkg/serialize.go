package pkg

import (
	"fmt"
	"strings"
)

type Foo struct {
	Name  string
	Space string
}

func Encode(f *Foo) string {
	return fmt.Sprintf("%s_%s", f.Name, f.Space)
}

func Decode(data string) *Foo {
	tmp := strings.Split(data, "_")
	return &Foo{
		Name:  tmp[0],
		Space: tmp[1],
	}
}
