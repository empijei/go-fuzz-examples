package simple

import "testing"

func TestReadNum(t *testing.T) {
	var tests = []struct {
		name string
		in   string
		want int
	}{
		{name: "valid", in: "42", want: 42},
		{name: "invalid", in: "foo", want: 0},
		{name: "negative", in: "-17", want: -17},
		{name: "zero", in: "0", want: 0},
		{name: "float", in: "1.7", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotn := ReadNum(tt.in)
			if gotn != tt.want {
				t.Errorf("ReadNum(%q): got: %d want: %d", tt.in, gotn, tt.want)
			}
			if gots := WriteNum(gotn); gotn != 0 && gots != tt.in {
				t.Errorf("WriteNum(%d): got: %q, want: %q", gotn, gots, tt.in)
			}
		})
	}
}
