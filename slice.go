package gogu

import (
	"errors"
	"fmt"
)

// Sum returns the sum of the slice items. These have to satisfy the type constraints declared as Number.
func Sum[T Number](slice []T) T {
	var acc T
	for _, v := range slice {
		acc += v
	}
	return acc
}

// SumBy is like Sum except it accept a callback function which is invoked
// for each element in the slice to generate the value to be summed.
func SumBy[T1 any, T2 Number](slice []T1, fn func(T1) T2) T2 {
	var acc T2
	for _, v := range slice {
		acc += fn(v)
	}
	return acc
}

// Mean computes the mean value of the slice elements.
func Mean[T Number](slice []T) T {
	var result T
	for i := 0; i < len(slice); i++ {
		result += slice[i]
	}
	return result / T(len(slice))
}

// IndexOf returns the index of the firs occurrence of a value
// in the slice, or -1 if value is not present in the slice.
func IndexOf[T comparable](s []T, val T) int {
	for k, v := range s {
		if v == val {
			return k
		}
	}

	return -1
}

// LastIndexOf returns the index of the last occurrence of a value.
func LastIndexOf[T comparable](s []T, val T) int {
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		if s[i] == val {
			return i
		}
	}
	return -1
}

// Map produces a new slice of values by mapping each value in the list through a transformation function.
func Map[T1, T2 any](slice []T1, fn func(T1) T2) []T2 {
	result := make([]T2, len(slice))

	for idx, v := range slice {
		result[idx] = fn(v)
	}
	return result
}

// ForEach iterates over the elements of a collection and invokes the callback fn function on each element.
func ForEach[T any](slice []T, fn func(T)) {
	for _, v := range slice {
		fn(v)
	}
}

// ForEachRight is the same as ForEach, but starts the iteration from the last element.
func ForEachRight[T any](slice []T, fn func(T)) {
	for i := len(slice) - 1; i >= 0; i-- {
		fn(slice[i])
	}
}

// Reduce reduces the collection to a value which is the accumulated result of running
// each element in the collection through the callback function yielding a single value.
func Reduce[T1, T2 any](slice []T1, fn func(T1, T2) T2, initVal T2) T2 {
	actual := initVal

	for _, v := range slice {
		actual = fn(v, actual)
	}

	return actual
}

// Reverse reverses the order of elements, so that the first element becomes the last,
// the second element becomes the second to last, and so on.
func Reverse[T any](sl []T) []T {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}

	return sl
}

// Unique returns the collection unique values.
func Unique[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	result := []T{}

	for _, v := range slice {
		if _, ok := keys[v]; !ok {
			keys[v] = true
			result = append(result, v)
		}
	}

	return result
}

// UniqueBy is like Unique except that it accept a callback function which is invoked on each
// element of the slice applying the criteria by which the uniqueness is computed.
func UniqueBy[T comparable](slice []T, fn func(T) T) []T {
	keys := make(map[T]bool)
	result := []T{}

	for _, v := range slice {
		if _, ok := keys[fn(v)]; !ok {
			keys[fn(v)] = true
			result = append(result, v)
		}
	}

	return result
}

// Every returns true if all the elements of a slice satisfies the criteria of the callback function.
func Every[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Some returns true if some elements of a slice satisfies the criteria of the callback function.
func Some[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}

// Partition splits the collection elements into two, the ones which satisfies the condition
// expressed in the callback function (fn) and those which does not satisfy the condition.
func Partition[T comparable](slice []T, fn func(T) bool) [2][]T {
	var result = [2][]T{}

	for _, v := range slice {
		if fn(v) {
			result[0] = append(result[0], v)
		} else {
			result[1] = append(result[1], v)
		}
	}

	return result
}

// Contains returns true if the value is present in the collection.
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Duplicate returns the duplicated values of a collection.
func Duplicate[T comparable](slice []T) []T {
	keyCount := make(map[T]int)
	result := make([]T, 0, len(slice))

	// Count how many times a value is showing up in the provided collection.
	for _, v := range slice {
		if _, ok := keyCount[v]; !ok {
			keyCount[v] = 1
		} else {
			keyCount[v]++
		}
	}

	// Include only the values which count frequency is greater than 1 into the resulting slice.
	for k, v := range keyCount {
		if v > 1 {
			result = append(result, k)
		}
	}
	return result
}

// DuplicateWithIndex puts the duplicated values of a collection into a map as a key value pair,
// where the key is the collection element and the value is its position.
func DuplicateWithIndex[T comparable](slice []T) map[T]int {
	var count int
	kvMap := make(map[T][]int)
	result := make(map[T]int)

	// Count how many times a value is showing up in the provided collection.
	for idx, v := range slice {
		if _, ok := kvMap[v]; !ok {
			// Create a slice with a dimension of 2, which first element contains the position (the index)
			// of the first found duplicate value and the second indicates the number of appearance.
			kvMap[v] = make([]int, 2)
			count = 1
			kvMap[v][0] = idx
			kvMap[v][1] = count
		} else {
			count++
			kvMap[v][1] = count
		}
	}

	// Include into the resulting slice only the values which frequency is greater than 1.
	for k, v := range kvMap {
		if v[1] > 1 {
			result[k] = v[0]
		}
	}
	return result
}

// Merge merges the first slice with the other slices defined as variadic parameter.
func Merge[T any](s []T, params ...[]T) []T {
	merged := make([]T, 0, len(s))

	for i := 0; i < len(params); i++ {
		merged = append(merged, params[i]...)
	}
	merged = append(s, merged...)

	return merged
}

// Flatten flattens the slice all the way down to the deepest nesting level.
func Flatten[T any](slice any) ([]T, error) {
	return baseFlatten([]T{}, slice)
}

func baseFlatten[T any](acc []T, slice any) ([]T, error) {
	var err error

	switch v := any(slice).(type) {
	case T:
		acc = append(acc, v)
	case []T:
		acc = append(acc, v...)
	case []any:
		for _, sv := range v {
			acc, err = baseFlatten(acc, sv)
			if err != nil {
				return nil, errors.New("flattening error")
			}
		}
	default:
		return nil, errors.New("flattening error")
	}

	return acc, nil
}

// Union computes the union of the passed\-in slice and returns an
// ordered list of unique items that are present in one or more of the slices.
func Union[T comparable](slice any) ([]T, error) {
	var err error
	if flatten, err := baseFlatten([]T{}, slice); err == nil {
		return Unique(flatten), nil
	}
	return nil, err
}

// Intersection computes the list of values that are the intersection of all the slices.
// Each value in the result should be present in each of the provided slices.
func Intersection[T comparable](params ...[]T) []T {
	result := []T{}

	for i := 0; i < len(params[0]); i++ {
		item := params[0][i]
		if Contains(result, item) {
			continue
		}
		var j int
		for j = 1; j < len(params); j++ {
			if !Contains(params[j], item) {
				break
			}
		}

		if j == len(params) {
			result = append(result, item)
		}
	}

	return result
}

// IntersectionBy is like Intersection, except that it accepts and callback function which is invoked on each element of the collection.
func IntersectionBy[T comparable](fn func(T) T, params ...[]T) []T {
	result := []T{}

	for i := 0; i < len(params[0]); i++ {
		item := params[0][i]
		if Contains(result, fn(item)) {
			continue
		}
		var j int
		for j = 1; j < len(params); j++ {
			has := func() bool {
				for _, v := range params[j] {
					if fn(v) == fn(item) {
						return true
					}
				}
				return false
			}
			if !has() {
				break
			}
		}

		if j == len(params) {
			result = append(result, item)
		}
	}

	return result
}

// Without returns a copy of the slice with all the values defined in the variadic parameter removed.
func Without[T1 comparable, T2 any](slice []T1, values ...T1) []T1 {
	keys := make(map[T1]bool)
	uni := make([]T1, 0, len(slice))
loop:
	for _, v := range slice {
		for _, val := range values {
			if v == val {
				continue loop
			}
		}
		if _, ok := keys[v]; !ok {
			keys[v] = true
			uni = append(uni, v)
		}
	}

	return uni
}

// Difference is similar to Without, but returns the values from
// the first slice that are not present in the second slice.
func Difference[T comparable](s1, s2 []T) []T {
	keys := make(map[T]bool)
	unique := []T{}
loop:
	for _, v := range s1 {
		for _, val := range s2 {
			if v == val {
				continue loop
			}
		}
		if _, ok := keys[v]; !ok {
			keys[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

// DifferenceBy is like Difference, except that invokes a callback function on each
// element of the slice, applying the criteria by which the difference is computed.
func DifferenceBy[T comparable](s1, s2 []T, fn func(T) T) []T {
	keys := make(map[T]bool)
	unique := []T{}
loop:
	for _, v := range s1 {
		for _, val := range s2 {
			if fn(v) == fn(val) {
				continue loop
			}
		}
		if _, ok := keys[v]; !ok {
			keys[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

// Chunk split the slice into groups of slices each having the length of size.
// In case the source slice cannot be distributed equally, the last slice will contain fewer elements.
func Chunk[T comparable](slice []T, size int) [][]T {
	var result = make([][]T, 0, len(slice)/2+1)

	if size <= 0 {
		panic("Chunk size should be greater than zero.")
	}
	for i := 0; i < len(slice); i++ {
		if i%size == 0 {
			if i+size < len(slice) {
				result = append(result, slice[i:i+size])
			} else {
				result = append(result, slice[i:])
			}
		}
	}
	return result
}

// Drop creates a new slice with n elements dropped from the beginning.
// If n < 0 the elements will be dropped from the back of the collection.
func Drop[T any](slice []T, n int) []T {
	if Abs(n) < len(slice) {
		if n > 0 {
			return slice[n:]
		} else {
			return slice[:len(slice)-Abs(n)]
		}
	}
	return []T{}
}

// DropWhile creates a new slice excluding the elements dropped from the beginning.
// Elements are dropped by applying the condition invoked in the callback function.
func DropWhile[T any](slice []T, fn func(T) bool) []T {
	result := make([]T, 0, len(slice))

	for _, v := range slice {
		if !fn(v) {
			result = append(result, v)
		}
	}

	return result
}

// DropRightWhile creates a new slice excluding the elements dropped from the end.
// Elements are dropped by applying the condition invoked in the callback function.
func DropRightWhile[T any](slice []T, fn func(T) bool) []T {
	result := make([]T, 0, len(slice))

	for i := len(slice) - 1; i >= 0; i-- {
		if !fn(slice[i]) {
			result = append(result, slice[i])
		}
	}

	return result
}

// MapByIndex
func mapByIndex[T1 comparable, T2 any](origSlice []T2, mapSlice []T1) map[T1][]T2 {
	result := make(map[T1][]T2)

	for idx, v := range mapSlice {
		if _, ok := result[v]; !ok {
			result[v] = make([]T2, 0, len(mapSlice))
		}
		result[v] = append(result[v], origSlice[idx])
	}

	return result
}

// GroupBy splits a collection into a key-value set, grouped by the result of running each value through the callback function fn.
// The return value is a map where the key is the conditional logic of the callback function
// and the values are the callback function returned values.
func GroupBy[T1, T2 comparable](slice []T1, fn func(T1) T2) map[T2][]T1 {
	return mapByIndex(slice, Map(slice, fn))
}

// Zip iteratively merges together the values of the slice parameters with the values at the corresponding position.
func Zip[T any](slices ...[]T) [][]T {
	var result = make([][]T, len(slices))
	var sliceLen int

	if len(slices) > 0 {
		sliceLen = len(slices[0])
	}

	if sliceLen != len(slices) {
		panic(fmt.Sprintf("the number of slice parameters (%d) does not match with the slice length (%d)", len(slices), sliceLen))
	}

	for idx, sl := range slices {
		if sliceLen != len(sl) {
			panic("the slice parameters should have identical length")
		}
		result[idx] = make([]T, len(sl))
	}

	for x := 0; x < sliceLen; x++ {
		for i := 0; i < len(slices); i++ {
			result[i][x] = slices[x][i]
		}
	}
	return result
}

// Unzip is the opposite of Zip: given a slice of slices it returns a series of new slices,
// the first of which contains all the first elements in the input slices,
// the second of which contains all the second elements, and so on.
func Unzip[T any](slices ...[]T) [][]T {
	var result = make([][]T, len(slices))
	var sliceLen int

	if len(slices) > 0 {
		sliceLen = len(slices[0])
	}

	if sliceLen != len(slices) {
		panic(fmt.Sprintf("the number of slice parameters (%d) does not match with the slice length (%d)", len(slices), sliceLen))
	}

	for idx, sl := range slices {
		if sliceLen != len(sl) {
			panic("the slice parameters should have identical length")
		}
		result[idx] = make([]T, len(sl))
	}

	for x := 0; x < sliceLen; x++ {
		for i := 0; i < len(slices); i++ {
			result[x][i] = slices[i][x]
		}
	}
	return result
}

// ToSlice returns the function arguments as a slice.
func ToSlice[T any](args ...T) []T {
	slice := make([]T, 0, len(args))
	slice = append(slice, args...)

	return slice
}

// ToMap convert slice to map
func ToMap[T, K comparable](slice []T, fn func(x T) K) map[K]T {
	result := make(map[K]T, len(slice))
	for _, v := range slice {
		k := fn(v)
		result[k] = v
	}
	return result
}
