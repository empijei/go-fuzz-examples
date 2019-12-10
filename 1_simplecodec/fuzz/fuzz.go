package simplecodec

import (
	"fmt"

	"github.com/empijei/gofuzz-talk/simple"
)

func Fuzz(b []byte) int {
	n := simple.ReadNum(string(b))
	if n == 0 {
		return 0
	}
	s := simple.WriteNum(n)
	if s != string(b) {
		fmt.Printf("got %q want %q\n", s, b)
		panic("encoding differs")
	}
	return 1
}
