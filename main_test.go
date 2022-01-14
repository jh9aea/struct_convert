package struct_convert

import (
	"testing"
)

type aStruct struct {
	A string
	B string
}

type bStruct struct {
	B string
}

func TestConvert(t *testing.T) {
	a := &aStruct{A: "a", B: "b"}
	var b bStruct
	var err error
	b, err = Convert[bStruct](a)
	if err != nil {
		t.Fatal(err)
	}
	if b.B != a.B {
		t.Fail()
	}
}
