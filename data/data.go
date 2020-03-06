package data

import (
	"reflect"
	"time"
)

func ShortContent(content interface{}, force ...string) (m map[string]interface{}) {
	m = make(map[string]interface{})
	var fields []string
	t := reflect.TypeOf(content)
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
	}
	val := reflect.ValueOf(content)
	var tt time.Time
	for _, fname := range fields {
		if val.FieldByName(fname).Interface() == "" ||
			val.FieldByName(fname).Interface() == 0 ||
			val.FieldByName(fname).Interface() == nil ||
			val.FieldByName(fname).Interface() == uint(0) ||
			val.FieldByName(fname).Interface() == int8(0) ||
			val.FieldByName(fname).Interface() == int64(0) ||
			val.FieldByName(fname).Interface() == float64(0) ||
			val.FieldByName(fname).Interface() == tt {
			if len(force) > 0 {
				for j := 0; j < len(force); j++ {
					if force[j] == fname {
						m[fname] = val.FieldByName(fname).Interface()
					}
				}
			}
		} else {
			// fmt.Printf("%s:%+v\r\n", fname, val.FieldByName(fname).Interface())
			m[fname] = val.FieldByName(fname).Interface()
		}
	}
	return
}
