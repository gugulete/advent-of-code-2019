package arrays

import "strconv"

//CopyIntArr copies array
func CopyIntArr(arr *[]int) []int {
	res := make([]int, len(*arr))
	for i, v := range *arr {
		res[i] = v
	}

	return res
}

// IntToStringArr converts int array to string array
func IntToStringArr(arr *[]int) []string {
	res := make([]string, len(*arr))
	for i, v := range *arr {
		res[i] = strconv.Itoa(v)
	}

	return res
}

// StringToIntArr converts string array to int array
func StringToIntArr(arr *[]string) []int {
	res := make([]int, len(*arr))
	for i, s := range *arr {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		res[i] = v
	}

	return res
}

// Index ...
func Index(vs []interface{}, t interface{}) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

// Includes ...
func Includes(vs []interface{}, t interface{}) bool {
	return Index(vs, t) >= 0
}

// Any ...
func Any(vs []interface{}, f func(interface{}) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All ...
func All(vs []interface{}, f func(interface{}) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter ...
func Filter(vs []interface{}, f func(interface{}) bool) []interface{} {
	vsf := make([]interface{}, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map ...
func Map(vs []interface{}, f func(interface{}) interface{}) []interface{} {
	vsm := make([]interface{}, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
