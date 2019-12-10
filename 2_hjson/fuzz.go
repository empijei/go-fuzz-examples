// hjson-go has 179 Stars, 21 Forks and it took 20 seconds to find 2 panics.
package hjson

import (
	"fmt"

	"github.com/dvyukov/go-fuzz-corpus/fuzz"
	hjson "github.com/hjson/hjson-go"
)

func Fuzz(data []byte) int {
	v0 := map[string]interface{}{}
	if hjson.Unmarshal(data, &v0) != nil {
		return 0
	}
	data1, err := hjson.Marshal(&v0)
	if err != nil {
		fmt.Printf("v0: %#v\n", v0)
		panic("can't marshal")
	}
	v1 := map[string]interface{}{}
	if err := hjson.Unmarshal(data1, &v1); err != nil {
		fmt.Printf("v0: %#v\n", v0)
		fmt.Printf("data1: %q\n", data1)
		panic("can't unmarshal")
	}
	if !fuzz.DeepEqual(v0, v1) {
		fmt.Printf("input: %q\n", data)
		fmt.Printf("v0: %#v\n", v0)
		fmt.Printf("v1: %#v\n", v1)
		panic("not equal")
	}
	return 1
}

/*
// Reference fuzzer taken from
// https://github.com/dvyukov/go-fuzz-corpus/blob/master/json/json.go

package json

import (
	"encoding/json"
	"fmt"

	"github.com/dvyukov/go-fuzz-corpus/fuzz"
)

func Fuzz(data []byte) int {
	score := 0
	for _, ctor := range []func() interface{}{
		func() interface{} { return nil },
		func() interface{} { return new([]interface{}) },
		func() interface{} { m := map[string]string{}; return &m },
		func() interface{} { m := map[string]interface{}{}; return &m },
		func() interface{} { return new(S) },
	} {
		v := ctor()
		if json.Unmarshal(data, v) != nil {
			continue
		}
		score = 1
		if s, ok := v.(*S); ok {
			if len(s.P) == 0 {
				s.P = []byte(`""`)
			}
		}
		data1, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		v1 := ctor()
		if json.Unmarshal(data1, v1) != nil {
			continue
		}
		if s, ok := v.(*S); ok {
			// Some additional escaping happens with P.
			s.P = nil
			v1.(*S).P = nil
		}
		if !fuzz.DeepEqual(v, v1) {
			fmt.Printf("v0: %#v\n", v)
			fmt.Printf("v1: %#v\n", v1)
			panic("not equal")
		}
	}
	return score
}

type S struct {
	A int    `json:",omitempty"`
	B string `json:"B1,omitempty"`
	C float64
	D bool
	E uint8
	F []byte
	G interface{}
	H map[string]interface{}
	I map[string]string
	J []interface{}
	K []string
	L S1
	M *S1
	N *int
	O **int
	P json.RawMessage
	Q Marshaller
	R int `json:"-"`
	S int `json:",string"`
}

type S1 struct {
	A int
	B string
}

type Marshaller struct {
	v string
}

func (m *Marshaller) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.v)
}

func (m *Marshaller) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &m.v)
}
*/
