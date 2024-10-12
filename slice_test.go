package gogu

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_Sum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, Sum([]int{1, 2, 3}))
	assert.Equal(12, SumBy([]int{1, 2, 3}, func(val int) int {
		return val * 2
	}))
	assert.Equal(6, SumBy([]string{"one", "two"}, func(val string) int {
		return len(val)
	}))
}

func TestSlice_Mean(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, Mean([]int{1, 2, 3}))
}

func TestFind_IndexOf(t *testing.T) {
	input := []int{1, 2, 3, 4, 2, -2, -1, 2}

	assert := assert.New(t)
	assert.Equal(0, IndexOf(input, 1))
	assert.Equal(-1, IndexOf(input, 5))
}

func TestFind_LastIndexOf(t *testing.T) {
	input := []int{1, 2, -1, 4, 2, -2, -1, 2}

	assert := assert.New(t)
	assert.Equal(6, LastIndexOf(input, -1))
	assert.Equal(-1, IndexOf(input, 5))
}

func TestSlice_Map(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{2, 4, 6}, Map([]int{1, 2, 3}, func(val int) int {
		return val * 2
	}))
	assert.Len(Map([]int{2, 4}, func(val int) int {
		return val * val
	}), 2)
}

func Example_map() {
	res := Map([]int{1, 2, 3}, func(val int) int {
		return val * 2
	})
	fmt.Println(res)

	// Output:
	// [2 4 6]
}

func TestSlice_ForEach(t *testing.T) {
	assert := assert.New(t)

	idx := 0
	input1 := []int{1, 2, 3, 4}
	output1 := make([]int, 4)

	ForEach(input1, func(val int) {
		output1[idx] = val
		idx++
	})
	assert.Equal(output1, input1)
	assert.IsIncreasing(output1)

	idx = 0
	input2 := []string{"a", "b", "c", "d"}
	output2 := make([]string, len(input2)-1)

	ForEach(input2, func(val string) {
		if idx != len(input1)-1 {
			output2[idx] = val
		}
		idx++
	})

	assert.Equal([]string{"a", "b", "c"}, output2)
	assert.Len(output2, 3)

	idx = 0
	ForEach(input2, func(val string) {
		input2[idx] = val + val
		idx++
	})
	assert.Equal([]string{"aa", "bb", "cc", "dd"}, input2)

	output3 := []string{}
	ForEachRight(input1, func(val int) {
		output3 = append(output3, strconv.Itoa(val))
	})
	assert.Equal([]string{"4", "3", "2", "1"}, output3)
}

func Example_forEach() {
	input := []int{1, 2, 3, 4}
	output := []int{}

	ForEach(input, func(val int) {
		val = val * 2
		output = append(output, val)
	})
	fmt.Println(output)

	// Output:
	// [2 4 6 8]
}

func TestSlice_Reduce(t *testing.T) {
	assert := assert.New(t)

	input1 := []int{1, 2, 3, 4}
	assert.Equal(10, Reduce(input1, func(a, b int) int {
		return a + b
	}, 0))

	input2 := []string{"a", "b", "c", "d"}
	assert.Equal("abcd", Reduce(input2, func(a, b string) string {
		return b + a
	}, ""))

	res := Reduce(input2, func(a, b string) string {
		return a + b
	}, "")
	res1 := []byte(res)
	sort.Slice(res1, func(i, j int) bool { return res[i] < res[j] })

	assert.Equal("abcd", string(res1))
}

func Example_reduce() {
	input1 := []int{1, 2, 3, 4}
	res1 := Reduce(input1, func(a, b int) int {
		return a + b
	}, 0)
	fmt.Println(res1)

	input2 := []string{"a", "b", "c", "d"}
	res2 := Reduce(input2, func(a, b string) string {
		return b + a
	}, "")
	fmt.Println(res2)

	// Output:
	// 10
	// abcd
}

func TestSlice_Reverse(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{4, 3, 2, 1}, Reverse([]int{1, 2, 3, 4}))
	assert.Equal([]string{"a", "b", "c"}, Reverse([]string{"c", "b", "a"}))

	assert.Equal("abcd", Reduce(Reverse([]string{"a", "b", "c", "d"}), func(a, b string) string {
		return a + b
	}, ""))

	assert.NotEqual("abcd", Reverse([]string{"a", "b", "c", "d"}))
}

func TestSlice_Unique(t *testing.T) {
	assert := assert.New(t)

	input := []int{1, 2, 4, 3, 1, 4, 5}
	res := Unique(input)

	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	assert.Equal([]int{1, 2, 3, 4, 5}, res)

	assert.Equal([]float64{2.1, 1.2}, UniqueBy([]float64{2.1, 1.2, 2.3}, func(v float64) float64 {
		return math.Floor(v)
	}))

	assert.Equal([]string{"a", "b", "c"}, UniqueBy([]string{"a", "b", "c", "B", "c", "A"}, func(v string) string {
		return strings.ToUpper(v)
	}))
}

func TestSlice_Every(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, Every([]int{2, 4, 6, 8, 10}, func(val int) bool {
		return val%2 == 0
	}))
	assert.NotEqual(true, Every([]int{-1, -2, 6, 8, 10}, func(val int) bool {
		return val > 0
	}))
	assert.Equal(false, Every([]any{"1", 1, 10, false}, func(val any) bool {
		return reflect.TypeOf(val).Kind() == reflect.Int
	}))
	assert.Equal(true, Every([]string{"1", "2", "3", "4"}, func(val string) bool {
		v, _ := strconv.Atoi(val)
		return reflect.TypeOf(v).Kind() == reflect.Int
	}))
}

func TestSlice_Some(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, Some([]int{1, 2, 3, 4, 5, 6}, func(val int) bool {
		return val%2 == 0
	}))
	assert.Equal(false, Some([]int{1, 3, 5}, func(val int) bool {
		return val%2 == 0
	}))
	assert.Equal(true, Some([]string{"1", "2", "3", "a"}, func(val string) bool {
		v, _ := strconv.Atoi(val)
		return reflect.TypeOf(v).Kind() == reflect.Int
	}))

	assert.Equal(false, Some([]string{"a", "b", "c"}, func(val string) bool {
		return reflect.TypeOf(val).Kind() == reflect.Int
	}))
}

func TestSlice_Partition(t *testing.T) {
	assert := assert.New(t)

	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Len(Partition(input, func(val int) bool {
		return val >= 5
	}), 2)

	res1 := Partition(input, func(val int) bool {
		return val < 5
	})
	assert.Equal([]int{0, 1, 2, 3, 4}, res1[0])

	res2 := Partition(input, func(val int) bool {
		return val < 0
	})
	assert.Empty(res2[0])
	assert.NotEmpty(res2[1])
}

func Example_partition() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res1 := Partition(input, func(val int) bool {
		return val >= 5
	})
	fmt.Println(res1)

	res2 := Partition(input, func(val int) bool {
		return val < 5
	})
	fmt.Println(res2)

	// Output:
	// [[5 6 7 8 9] [0 1 2 3 4]]
	// [[0 1 2 3 4] [5 6 7 8 9]]
}

func TestSlice_Contains(t *testing.T) {
	assert := assert.New(t)

	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Empty(Contains(input, -1))
	assert.Equal(true, Contains(input, 0))
	assert.NotEqual(true, Contains(input, 100))
}

func TestSlice_Duplicate(t *testing.T) {
	assert := assert.New(t)

	input1 := []int{-1, -1, 0, 1, 2, 3, 2, 5, 1, 6}
	assert.NotEmpty(Duplicate(input1))
	assert.Len(Duplicate(input1), 3)
	assert.ElementsMatch([]int{-1, 1, 2}, Duplicate(input1))

	input2 := []string{"One", "Two", "Three", "two", "One"}
	assert.ElementsMatch([]string{"One"}, Duplicate(input2))
	assert.ElementsMatch([]string{"one", "two"}, Duplicate(Map(input2, func(val string) string {
		return strings.ToLower(val)
	})))

	assert.Len(DuplicateWithIndex(input1), 3)
	res := DuplicateWithIndex(input1)

	indices := make([]int, 0, len(input1))
	for k := range res {
		indices = append(indices, k)
	}
	assert.ElementsMatch([]int{-1, 1, 2}, indices)
}

func Example_duplicate() {
	input1 := []int{-1, -1, 0, 1, 2, 3, 2, 5, 1, 6}
	res1 := Duplicate(input1)
	sort.Slice(res1, func(a, b int) bool { return res1[a] < res1[b] })
	fmt.Println(res1)

	res2 := DuplicateWithIndex(input1)
	fmt.Println(res2)

	// Output:
	// [-1 1 2]
	// map[-1:0 1:3 2:4]
}

func TestSlice_Merge(t *testing.T) {
	assert := assert.New(t)

	sl1 := []int{1, 2, 3, 4}
	sl2 := []int{5, 6, 7, 8}

	assert.Len(Merge(sl1, sl2), len(sl1)+len(sl2))
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8}, Merge(sl1, sl2))
	assert.Equal([]int{1, 2, 3}, Merge([]int{1}, []int{2}, []int{3}))
}

func TestSlice_Flatten(t *testing.T) {
	assert := assert.New(t)

	input1 := []any{[]float64{1.0, 2.0}, 1.1}
	result1, err := Flatten[float64](input1)
	assert.Equal([]float64{1.0, 2.0, 1.1}, result1)
	assert.NotNil(result1)
	assert.NoError(err)
	assert.Len(result1, 3)

	input2 := []any{[]float32{1.0, 2.0}, 3.0}
	result2, err := Flatten[float32](input2)
	assert.Error(err)
	assert.Nil(result2) // result is nil, because the last element in the slice is of type float64

	input3 := []string{"a", "b", "c"}
	result3, err := Flatten[string](Merge(input3, []string{"d", "e"}))
	assert.Equal([]string{"a", "b", "c", "d", "e"}, result3)
	assert.NotNil(result3)
	assert.NoError(err)

	input4 := []any{[]int{1, 2, 3}, []any{[]int{4}, 5}}
	result4, _ := Flatten[int](input4)
	assert.Equal([]int{1, 2, 3, 4, 5}, result4)

	int1 := []int{1, 2}
	res1 := Map([]string{"3", "4"}, func(val string) int {
		res, _ := strconv.Atoi(val)
		return res
	})
	result5, err := Flatten[int]([]any{int1, res1})
	assert.Equal([]int{1, 2, 3, 4}, result5)
	assert.NoError(err)
}

func Example_flatten() {
	input := []any{[]int{1, 2, 3}, []any{[]int{4}, 5}}
	result, _ := Flatten[int](input)
	fmt.Println(result)

	// Output:
	// [1 2 3 4 5]
}

func TestSlice_Union(t *testing.T) {
	assert := assert.New(t)

	input1 := []any{[]any{1, 2, []any{3, []int{4, 5, 6}}}, 7, []int{1, 2}, 3, []int{4, 7}, 8, 9, 9}
	result1, err := Union[int](input1)
	assert.Len(result1, 9)
	assert.NoError(err)
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, result1)

	input2 := []any{[]any{"One", "Two", []any{"Foo", []string{"Bar", "Baz", "Qux"}}}, "Foo", []string{"Foo", "Two"}, "Baz", "bar"}
	result2, err := Union[string](input2)
	assert.Len(result2, 7)
	assert.NoError(err)
	assert.Equal([]string{"One", "Two", "Foo", "Bar", "Baz", "Qux", "bar"}, result2)

	resMap := Map(result2, func(val string) string {
		return strings.ToLower(val)
	})
	result3, err := Union[string](resMap)
	assert.Equal([]string{"one", "two", "foo", "bar", "baz", "qux"}, result3)
	assert.NoError(err)
}

func Example_union() {
	input := []any{[]any{1, 2, []any{3, []int{4, 5, 6}}}, 7, []int{1, 2}, 3, []int{4, 7}, 8, 9, 9}
	res, _ := Union[int](input)
	fmt.Println(res)

	// Output:
	// [1 2 3 4 5 6 7 8 9]
}

func TestSlice_Intersection(t *testing.T) {
	assert := assert.New(t)

	result1 := Intersection([]int{1, 2, 4}, []int{0, 2, 1}, []int{2, 1, -2})
	assert.Equal([]int{1, 2}, result1)

	result2 := Intersection([]int{-1, 0}, []int{2, 3})
	assert.Empty(result2)
	assert.Equal([]int{}, result2)

	result3 := Intersection([]int{0, 1, 2}, []int{2, 0, 1}, []int{2, 1, 0})
	assert.Equal([]int{0, 1, 2}, result3)

	result4 := Intersection([]string{"a", "b"}, []string{"a", "a", "a"}, []string{"b", "a", "e"})
	assert.Equal([]string{"a"}, result4)
}

func Example_intersection() {
	res1 := Intersection([]int{1, 2, 4}, []int{0, 2, 1}, []int{2, 1, -2})
	fmt.Println(res1)

	res2 := Intersection([]string{"a", "b"}, []string{"a", "a", "a"}, []string{"b", "a", "e"})
	fmt.Println(res2)

	// Output:
	// [1 2]
	// [a]
}

func TestSlice_IntersectionBy(t *testing.T) {
	assert := assert.New(t)

	result1 := IntersectionBy(func(v float64) float64 {
		return math.Floor(v)
	}, []float64{2.1, 1.2}, []float64{2.3, 3.4}, []float64{1.0, 2.3})
	assert.Equal([]float64{2.1}, result1)

	result2 := IntersectionBy(func(v int) int {
		return v % 2
	}, []int{1, 2}, []int{2, 1})
	assert.Equal([]int{1, 2}, result2)

	result3 := IntersectionBy(func(v float64) float64 {
		return math.Floor(v)
	}, []float64{1.1, 2.0, 3.2}, []float64{4.0})
	assert.Equal([]float64{}, result3)
}

func Example_intersectioBy() {
	result1 := IntersectionBy(func(v float64) float64 {
		return math.Floor(v)
	}, []float64{2.1, 1.2}, []float64{2.3, 3.4}, []float64{1.0, 2.3})
	fmt.Println(result1)

	result2 := IntersectionBy(func(v int) int {
		return v % 2
	}, []int{1, 2}, []int{2, 1})
	fmt.Println(result2)

	// Output:
	// [2.1]
	// [1 2]
}

func TestSlice_Without(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{3}, Without[int, int]([]int{2, 1, 2, 3}, 1, 2))
	assert.Equal([]int{1, 2}, Without[int, int]([]int{1, 2, 3, 4}, 3, 4))
	assert.Equal([]int{1, 2}, Without[int, int]([]int{0, 1, 2, 3, 4, 5}, 0, 3, 4, 5))
	assert.Equal([]float64{1.0, 2.2}, Without[float64, float64]([]float64{1.0, 2.2, 3.0, 4.2}, 3.0, 4.2))

	assert.Empty(Without[int, int]([]int{}, 1, 2, 3, 4))
	assert.Empty(Without[int, int]([]int{0, 1, 2}, 0, 1, 2))
	assert.Empty(Without[int, int]([]int{}, 0, 1, 2))
	assert.Empty(Without[int, int]([]int{}))
}

func Example_without() {
	fmt.Println(Without[int, int]([]int{2, 1, 2, 3}, 1, 2))
	fmt.Println(Without[int, int]([]int{1, 2, 3, 4}, 3, 4))
	fmt.Println(Without[int, int]([]int{0, 1, 2, 3, 4, 5}, 0, 3, 4, 5))
	fmt.Println(Without[float64, float64]([]float64{1.0, 2.2, 3.0, 4.2}, 3.0, 4.2))

	// Output:
	// [3]
	// [1 2]
	// [1 2]
	// [1 2.2]
}

func TestSlice_Difference(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{1, 3}, Difference([]int{1, 2, 3, 4}, []int{2, 4}))
	assert.Equal([]int{1, 2, 3}, Difference([]int{1, 2, 3, 4}, []int{4, 5, 6, 7}))
	assert.Equal([]int{1, 2, 3, 4}, Difference([]int{1, 2, 3, 4}, []int{}))
	assert.Equal([]int{}, Difference([]int{1}, []int{1}))

	assert.Empty(Difference([]int{}, []int{1, 2, 3, 4}))
	assert.Empty(Difference([]int{}, []int{-1}))

	assert.Equal([]float64{1.2}, DifferenceBy([]float64{2.1, 1.2}, []float64{2.3, 3.4}, func(v float64) float64 {
		return math.Floor(v)
	}))
	assert.Equal([]float64{}, DifferenceBy([]float64{1.2}, []float64{1.4}, func(v float64) float64 {
		return math.Floor(v)
	}))
	assert.Equal([]int{1}, DifferenceBy([]int{1}, []int{4}, func(v int) int {
		return v % 2
	}))
	assert.Empty(DifferenceBy([]int{2}, []int{4}, func(v int) int {
		return v % 2
	}))
}

func TestSlice_Chunk(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{{0, 1}, {2, 3}}, Chunk([]int{0, 1, 2, 3}, 2))
	assert.Equal([][]int{{0, 1}, {2, 3}, {4}}, Chunk([]int{0, 1, 2, 3, 4}, 2))
	assert.Equal([][]int{{0}, {1}}, Chunk([]int{0, 1}, 1))
	assert.Equal([][]string{{"Tyrone", "Elie"}, {"Aidan", "Sam"}, {"Little Timmy"}}, Chunk([]string{"Tyrone", "Elie", "Aidan", "Sam", "Little Timmy"}, 2))

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	assert.Equal([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12}}, Chunk(input, 5))
	assert.Len(Chunk(input, 5), 3)
	assert.Len(Chunk(input, 12), 1)
	assert.Panics(func() { Chunk([]int{0, 1}, 0) })
}

func Example_chunk() {
	fmt.Println(Chunk([]int{0, 1, 2, 3}, 2))
	fmt.Println(Chunk([]int{0, 1, 2, 3, 4}, 2))
	fmt.Println(Chunk([]int{0, 1}, 1))

	// Output:
	// [[0 1] [2 3]]
	// [[0 1] [2 3] [4]]
	// [[0] [1]]
}

func TestSlice_Drop(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{1, 2}, Drop([]int{0, 1, 2}, 1))
	assert.Equal([]int{2}, Drop([]int{0, 1, 2}, 2))
	assert.Equal([]int{0, 1, 2}, Drop([]int{0, 1, 2}, 0))
	assert.Equal([]int{}, Drop([]int{0, 1, 2}, 3))
	assert.Equal([]int{}, Drop([]int{0, 1, 2}, 4))

	assert.Equal([]int{0, 1}, Drop([]int{0, 1, 2}, -1))
	assert.Equal([]int{}, Drop([]int{0, 1, 2}, -3))
	assert.Equal([]int{}, Drop([]int{0, 1, 2}, -4))

	assert.Equal([]string{"a", "AA"}, DropWhile([]string{"a", "AA", "bbb", "ccc"}, func(elem string) bool {
		return len(elem) > 2
	}))
	assert.Equal([]int{1, 3}, DropWhile([]int{1, 2, 3, 4}, func(val int) bool {
		return val%2 == 0
	}))
	assert.Equal([]int{2}, DropWhile([]int{1, 2}, func(val int) bool {
		return val < 2
	}))
	assert.Empty(DropWhile([]int{1, 2}, func(val int) bool {
		return val <= 2
	}))

	assert.Equal([]string{"AA", "a"}, DropRightWhile([]string{"a", "AA", "bbb", "ccc"}, func(elem string) bool {
		return len(elem) > 2
	}))
	assert.Equal([]int{3, 1}, DropRightWhile([]int{1, 2, 3, 4}, func(val int) bool {
		return val%2 == 0
	}))
	assert.Equal([]int{2}, DropRightWhile([]int{1, 2}, func(val int) bool {
		return val < 2
	}))
	assert.Empty(DropRightWhile([]int{1, 2}, func(val int) bool {
		return val <= 2
	}))
}

func Example_drop() {
	res := DropWhile([]string{"a", "aa", "bbb", "ccc"}, func(elem string) bool {
		return len(elem) > 2
	})
	fmt.Println(res)

	// Output:
	// [a aa]
}

func TestSlice_GroupBy(t *testing.T) {
	assert := assert.New(t)

	input1 := []float64{1.3, 1.5, 2.1, 2.9}
	assert.Len(GroupBy(input1, func(val float64) float64 {
		return math.Floor(val)
	}), 2)
	assert.Equal(map[float64][]float64{1: {1.3, 1.5}, 2: {2.1, 2.9}}, GroupBy(input1, func(val float64) float64 {
		return math.Floor(val)
	}))
	assert.Equal(map[float64][]float64{1: {1.3, 1.5}, 2: {2.1, 2.9}}, GroupBy(input1, func(val float64) float64 {
		return math.Floor(val)
	}))
	res1 := GroupBy(input1, func(val float64) int {
		if math.Floor(val) == 1 {
			return 2
		}
		return int(math.Floor(val))
	})
	assert.Len(res1, 1)
	assert.Equal(map[int][]float64{2: {1.3, 1.5, 2.1, 2.9}}, res1)

	input2 := []string{"one", "two", "three"}
	assert.Equal(map[int][]string{3: {"one", "two"}, 5: {"three"}}, GroupBy(input2, func(val string) int {
		return len(val)
	}))

	input3 := []string{"1", "2", "3"}
	assert.Len(GroupBy(input3, func(val string) int {
		res, _ := strconv.Atoi(val)
		return res
	}), 3)

	res3 := GroupBy(input3, func(val string) int {
		res, _ := strconv.Atoi(val)
		return res
	})
	for k, v := range res3 {
		val, _ := strconv.Atoi(v[0])
		assert.Equal(k, val)
	}
}

func Example_groupBy() {
	input := []float64{1.3, 1.5, 2.1, 2.9}
	res := GroupBy(input, func(val float64) float64 {
		return math.Floor(val)
	})
	fmt.Println(res)

	// Output:
	// map[1:[1.3 1.5] 2:[2.1 2.9]]
}

func TestSlice_ToSlice(t *testing.T) {
	assert := assert.New(t)

	assert.Len(ToSlice[int](), 0)
	assert.Empty(ToSlice[string]())
	assert.Equal([]int{1, 2}, ToSlice(1, 2))
	assert.Equal([]int{1, 2, 3}, ToSlice(1, 2, 3))
	assert.Equal([]string{"a", "b"}, ToSlice("a", "b"))
}

func TestSlice_Zip(t *testing.T) {
	assert := assert.New(t)

	input := [][]any{{"one", "two"}, {1, 2}}
	assert.Panics(func() { Zip(input) })
	assert.Equal([][]any{{"one", 1}, {"two", 2}}, Zip([]any{"one", "two"}, []any{1, 2}))
}

func TestSlice_Unzip(t *testing.T) {
	assert := assert.New(t)

	res := Unzip([]any{"one", 1}, []any{"two", 2})
	assert.Equal([][]any{{"one", "two"}, {1, 2}}, res)
}

func Example_zip() {
	res := Zip([]any{"one", "two"}, []any{1, 2})
	fmt.Println(res)

	// Output:
	// [[one 1] [two 2]]
}

func Example_unzip() {
	res := Unzip([]any{"one", 1}, []any{"two", 2})
	fmt.Println(res)

	// Output:
	// [[one two] [1 2]]
}

func TestSlice_ToMap(t *testing.T) {
	type some_struct struct {
		Id    uint
		Title string
	}

	s1 := some_struct{Id: 1, Title: "foo"}
	s2 := some_struct{Id: 10, Title: "bar"}
	s3 := some_struct{Id: 100, Title: "baz"}

	input := []some_struct{
		s1,
		s2,
		s3,
	}

	result := ToMap(input, func(x some_struct) uint { return x.Id })

	assert := assert.New(t)
	assert.True(len(result) == 3)
	assert.Equal(map[uint]some_struct{s1.Id: s1, s2.Id: s2, s3.Id: s3}, result)
}

func Example_toMap() {
	slice := []int{1, 2, 3}
	result := ToMap(slice, func(x int) int { return x })

	fmt.Println(result)
	// {1: 1, 2: 2, 3: 3}
}
