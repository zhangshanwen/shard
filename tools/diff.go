package tools

import (
	"fmt"
	"reflect"
)

func DiffStruct(origin, compare interface{}, tag string) (diff string) {
	ot := reflect.TypeOf(origin)
	ov := reflect.ValueOf(origin)
	ct := reflect.TypeOf(compare)
	cv := reflect.ValueOf(compare)

	for k := 0; k < ot.NumField(); k++ {
		for j := 0; j < ct.NumField(); j++ {
			if ot.Field(k).Name == ct.Field(j).Name && ov.Field(k).Interface() != cv.Field(j).Interface() {
				key := ot.Field(k).Tag.Get(tag)
				diff = fmt.Sprintf("%v:%v->%v ", key, cv.Field(j).Interface(), ov.Field(k).Interface())
			}
		}
	}
	return
}
func DiffStructIsChange(origin, compare interface{}) (isChange bool) {
	ot := reflect.TypeOf(origin)
	ov := reflect.ValueOf(origin)
	ct := reflect.TypeOf(compare)
	cv := reflect.ValueOf(compare)

	for k := 0; k < ot.NumField(); k++ {
		for j := 0; j < ct.NumField(); j++ {
			if ot.Field(k).Name == ct.Field(j).Name && ov.Field(k).Interface() != cv.Field(j).Interface() {
				return true
			}
		}
	}
	return
}
