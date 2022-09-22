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

func Diff2StructMap(origin, compare interface{}, tag string) (change map[string]interface{}) {
	change = make(map[string]interface{})
	ot := reflect.TypeOf(origin)
	ov := reflect.ValueOf(origin)
	ct := reflect.TypeOf(compare)
	cv := reflect.ValueOf(compare)

	for k := 0; k < ot.NumField(); k++ {
		for j := 0; j < ct.NumField(); j++ {
			if ot.Field(k).Name == ct.Field(j).Name && ov.Field(k).Interface() != cv.Field(j).Interface() {
				key := ot.Field(k).Tag.Get(tag)
				change[key] = cv.Field(j).Interface()
			}
		}
	}
	return
}

func Diff2StructMapAndFromTo(origin, compare interface{}, tag string) (change, fromTo map[string]interface{}) {
	change = make(map[string]interface{})
	fromTo = make(map[string]interface{})
	ot := reflect.TypeOf(origin)
	ov := reflect.ValueOf(origin)
	ct := reflect.TypeOf(compare)
	cv := reflect.ValueOf(compare)

	for k := 0; k < ot.NumField(); k++ {
		for j := 0; j < ct.NumField(); j++ {
			if ot.Field(k).Name == ct.Field(j).Name && ov.Field(k).Interface() != cv.Field(j).Interface() {
				key := ot.Field(k).Tag.Get(tag)
				change[key] = cv.Field(j).Interface()
				fromTo[key] = map[string]interface{}{"to": cv.Field(j).Interface(), "from": ov.Field(k).Interface()}
			}
		}
	}
	return
}
