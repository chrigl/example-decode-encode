package pkg

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
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
		{
			input: Foo{
				Name:     "hello",
				Space:    "world",
				NewField: "bar",
			},
			expect: "hello_world_bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.expect, func(t *testing.T) {
			got := Encode(&tt.input)
			if got != tt.expect {
				t.Errorf("want: %s, got: %s", tt.expect, got)
			}

			back := Decode(got)
			if *back != tt.input {
				t.Errorf("unable to Decode back. want: %+v, got: %+v", tt.input, back)
			}
		})
	}
}
