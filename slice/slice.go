package slice

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
func RemoveElementInt(arr []int, elem int) []int {
	if len(arr) == 0 {
		return arr
	}
	for i, v := range arr {
		if v == elem {
			arr = append(arr[:i], arr[i+1:]...)
			return RemoveElementInt(arr, elem)
		}
	}
	return arr
}
