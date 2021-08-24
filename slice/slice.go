package slice

import (
	"math"
	"strings"
	"time"
)

func FindElement(arr []interface{}) (newArr []interface{}) {
	newArr = make([]interface{}, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func RemoveRepeatedElement(arr []interface{}) (newArr []interface{}) {
	newArr = make([]interface{}, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func RemoveRepeatedString(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
func RemoveRepeatedInt(arr []int) (newArr []int) {
	newArr = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func RemoveRepeatedFloat64(arr []float64) (newArr []float64) {
	newArr = make([]float64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func RemoveElement(arr []interface{}, elem interface{}) []interface{} {
	if len(arr) == 0 {
		return arr
	}
	for i, v := range arr {
		if v == elem {
			arr = append(arr[:i], arr[i+1:]...)
			return RemoveElement(arr, elem)
		}
	}
	return arr
}
func RemoveInt(arr []int, elem int) []int {
	if len(arr) == 0 {
		return arr
	}
	for i, v := range arr {
		if v == elem {
			arr = append(arr[:i], arr[i+1:]...)
			return RemoveInt(arr, elem)
		}
	}
	return arr
}
func RemoveString(arr []string, elem string) []string {
	if len(arr) == 0 {
		return arr
	}
	for i, v := range arr {
		if v == elem {
			arr = append(arr[:i], arr[i+1:]...)
			return RemoveString(arr, elem)
		}
	}
	return arr
}
func RemoveZero(slice []interface{}) []interface{} {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if IfZero(v) {
			slice = append(slice[:i], slice[i+1:]...)
			return RemoveZero(slice)
		}
	}
	return slice
}

//判断一个值是否为零值，只支持string,float,int,time 以及其各自的指针，"%"和"%%"也属于零值范畴，场景是like语句
func IfZero(arg interface{}) bool {
	if arg == nil {
		return true
	}
	switch v := arg.(type) {
	case int, int32, int16, int64:
		if v == 0 {
			return true
		}
	case float32:
		r := float64(v)
		return math.Abs(r-0) < 0.0000001
	case float64:
		return math.Abs(v-0) < 0.0000001
	case string:
		if v == "" || v == "%%" || v == "%" {
			return true
		}
	case *string, *int, *int64, *int32, *int16, *int8, *float32, *float64, *time.Time:
		if v == nil {
			return true
		}
	case time.Time:
		return v.IsZero()
	default:
		return false
	}
	return false
}

//去重合并
func FindAddString(old, now []string) (diff []string) {
	for i := 0; i < len(old); i++ {
		for j := 0; j < len(now); j++ {
			if old[i] == now[j] {
				break
			}
		}
		diff = append(diff, old[i])
	}
	return
}

//指定删除
func FindSubString(old []string, now string) (diff []string) {
	for i := 0; i < len(old); i++ {
		if strings.Contains(now, old[i]) {

		} else {
			diff = append(diff, old[i])
		}
	}
	return
}

//指定查找
func FindInt(old []int, now int) bool {
	for i := 0; i < len(old); i++ {
		if now == old[i] {
			return true
		}
	}
	return false
}
