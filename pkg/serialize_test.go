package pkg

import (
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		input  Foo
		expect string
	}{
		{
			input: Foo{
				Name:  "hello",
				Space: "world",
			},
			expect: "hello_world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.expect, func(t *testing.T) {
			got := Encode(&tt.input)
			if got != tt.expect {
				t.Errorf("want: %s, got: %s", tt.expect, got)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		input  string
		expect Foo
	}{
		{
			input: "hello_world",
			expect: Foo{
				Name:  "hello",
				Space: "world",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := Decode(tt.input)
			if *got != tt.expect {
				t.Errorf("want: %+v, go: %+v", tt.expect, got)
			}
		})
	}
}
