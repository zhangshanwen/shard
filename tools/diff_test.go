package tools

import (
	"reflect"
	"testing"
)

type (
	A struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}
)

func TestDiff2StructMap(t *testing.T) {
	a := A{Id: 1, Name: "李四"}
	b := A{Id: 1, Name: "李五"}
	t.Log(reflect.TypeOf(a).Name())
	t.Log(Diff2StructMap(a, b, "json"))
	t.Log(Diff2StructMapAndFromTo(a, b, "json"))

}
