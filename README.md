# gogu (Go Generics Utility)
![Coverage](https://img.shields.io/badge/Coverage-93.8%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/esimov/gogu)](https://goreportcard.com/report/github.com/esimov/gogu)
[![CI](https://github.com/esimov/gogu/actions/workflows/ci.yml/badge.svg)](https://github.com/esimov/gogu/actions/workflows/ci.yml)
[![Go Reference](https://img.shields.io/badge/pkg.go.dev-reference-007d9c?logo=go)](https://pkg.go.dev/github.com/esimov/gogu)
[![release](https://img.shields.io/badge/Release-v1.0.3-blue.svg)](https://github.com/esimov/gogu/releases/tag/v1.0.3)
[![license](https://img.shields.io/github/license/esimov/pigo)](./LICENSE)

**Gogu** is a versatile, comprehensive, reusable and efficient concurrent-safe utility functions and data structures library taking advantage of the Go generics. It was inspired by other well known and established frameworks like [lodash](https://lodash.com/) or [Apache Commons](https://commons.apache.org/) and some concepts being more closer to the functional programming paradigms.

Its main purpose is to facilitate the ease of working with common data structures like slices, maps and strings, through the implementation of many utility functions commonly used in the day-by-day jobs, but also integrating some of the most used data structure algorithms.

## ✨ Features
**In what's different this library from other Go libraries exploring Go generics?** 
- [x] It's concurrent-safe (with the exception of B-tree package)
- [x] Implements a dozens of time related functions like: [`before`](<#func-before>), [`after`](<#func-after>), [`delay`](<#func-delay>), [`memoize`](<#func-memoizert-v-memoize>), [`debounce`](<#func-newdebounce>), [`once`](<#func-once>), [`retry`](<#func-rtypet-retry>)
- [x] Rich utility functions to operate with strings
- [x] Very wide range of supported functions to deal with slice and map operations
- [x] Extensive test coverage
- [x] Implements the most used data structures
- [x] Thourough documentation accompanied with examples

## 🚀 Run

```bash
$ go get github.com/esimov/gogu
```

## 🛠 Usage
```go
package main

import "github.com/esimov/gogu"

func main() {
  // main program
}
```

## 📖 Specifications
- **Generic Data Structures**
  - [`bst`](https://github.com/esimov/gogu/tree/master/bstree): Binary Search Tree data structure implementation, where each node has at most two child nodes and the key of its internal node is greater than all the keys in the respective node's left subtree and less than the ones in the right subtree
  - [`btree`](https://github.com/esimov/gogu/tree/master/btree): B-tree data structure implementation which is a self-balancing tree data structure maintaining its values in sorted order
  - [`cache`](https://github.com/esimov/gogu/tree/master/cache): a basic in-memory key-value storage system
  - [`heap`](https://github.com/esimov/gogu/tree/master/heap): Binary Heap data structure implementation where each node of the subtree is greather or equal then the parent node
  - [`list`](https://github.com/esimov/gogu/tree/master/list): implements a singly and doubly linked list data structure
  - [`queue`](https://github.com/esimov/gogu/tree/master/queue): package queue implements a FIFO (First-In-First-Out) data structure in two forms: using as storage system a resizing array and a doubly linked list
  - [`stack`](https://github.com/esimov/gogu/tree/master/stack): package stack implements a LIFO (Last-In-First-Out) data structure where the last element added to the stack is processed first
  - [`trie`](https://github.com/esimov/gogu/tree/master/trie): package trie provides a thread safe implementation of the ternary search tree data structure. Tries are used for locating specific keys from within a set or for quick lookup searches within a text like auto-completion or spell checking.

- **General utility functions**
  - [Abs](<#func-abs>)
  - [Clamp](<#func-clamp>)
  - [Compare](<#func-compare>)
  - [Equal](<#func-equal>)
  - [InRange](<#func-inrange>)
  - [Invert](<#func-invert>)
  - [Less](<#func-less>)
  - [Max](<#func-max>)
  - [Min](<#func-min>)

- **Strings utility functions**
  - [CamelCase](<#func-camelcase>)
  - [Capitalize](<#func-capitalize>)
  - [KebabCase](<#func-kebabcase>)
  - [N](<#func-n>)
  - [Null](<#func-null>)
  - [NumToString](<#func-numtostring>)
  - [Pad](<#func-pad>)
  - [PadLeft](<#func-padleft>)
  - [PadRight](<#func-padright>)
  - [ReverseStr](<#func-reversestr>)
  - [SnakeCase](<#func-snakecase>)
  - [SplitAtIndex](<#func-splitatindex>)
  - [Substr](<#func-substr>)
  - [ToLower](<#func-tolower>)
  - [ToUpper](<#func-toupper>)
  - [Unwrap](<#func-unwrap>)
  - [Wrap](<#func-wrap>)
  - [WrapAllRune](<#func-wrapallrune>)

- **Slice utility functions**
  - [Chunk](<#func-chunk>)
  - [SumBy](#sumby)
  - [Contains](<#func-contains>)
  - [Difference](<#func-difference>)
  - [DifferenceBy](<#func-differenceby>)
  - [Drop](<#func-drop>)
  - [DropWhile](<#func-dropwhile>)
  - [DropRightWhile](<#func-droprightwhile>)
  - [Duplicate](<#func-duplicate>)
  - [DuplicateWithIndex](<#func-duplicatewithindex>)
  - [Every](<#func-every>)
  - [Filter](<#func-filter>)
  - [FindAll](<#func-findall>)
  - [FindIndex](<#func-findindex>)
  - [FindLastIndex](<#func-findlastindex>)
  - [FindMax](<#func-findmax>)
  - [FindMaxBy](<#func-findmaxby>)
  - [FindMaxByKey](<#func-findmaxbykey>)
  - [FindMin](<#func-findmin>)
  - [FindMinBy](<#func-findminby>)
  - [FindMinByKey](<#func-findminbykey>)
  - [Flatten](<#func-flatten>)
  - [ForEach](<#func-foreach>)
  - [ForEachRight](<#func-foreachright>)
  - [GroupBy](<#func-groupby>)
  - [IndexOf](<#func-indexof>)
  - [Intersection](<#func-intersection>)
  - [IntersectionBy](<#func-intersectionby>)
  - [LastIndexOf](<#func-lastindexof>)
  - [Mean](<#func-mean>)
  - [Merge](<#func-merge>)
  - [Nth](<#func-nth>)
  - [Partition](<#func-partition>)
  - [PartitionMap](<#func-partitionmap>)
  - [Pluck](<#func-pluck>)
  - [Range](<#func-range>)
  - [RangeRight](<#func-rangeright>)
  - [Reduce](<#func-reduce>)
  - [Reject](<#func-reject>)
  - [Reverse](<#func-reverse>)
  - [Shuffle](<#func-shuffle>)
  - [SliceToMap](<#func-slicetomap>)
  - [Some](<#func-some>)
  - [Sum](<#func-sum>)
  - [SumBy](<#func-sumby>)
  - [ToSlice](<#func-toslice>)
  - [Union](<#func-union>)
  - [Unique](<#func-unique>)
  - [UniqueBy](<#func-uniqueby>)
  - [Unzip](<#func-unzip>)
  - [Without](<#func-without>)
  - [Zip](<#func-zip>)

- **Map utility functions**
  - [Filter2DMapCollection](<#func-filter2dmapcollection>)
  - [FilterMap](<#func-filtermap>)
  - [FilterMapCollection](<#func-filtermapcollection>)
  - [Find](<#func-find>)
  - [FindByKey](<#func-findbykey>)
  - [FindKey](<#func-findkey>)
  - [Keys](<#func-keys>)
  - [Map](<#func-map>)
  - [MapCollection](<#func-mapcollection>)
  - [MapContains](<#func-mapcontains>)
  - [MapEvery) bool](<#func-mapevery>)
  - [MapKeys](<#func-mapkeys>)
  - [MapSome](<#func-mapsome>)
  - [MapUnique](<#func-mapunique>)
  - [MapValues](<#func-mapvalues>)
  - [Omit](<#func-omit>)
  - [OmitBy](<#func-omitby>)
  - [Pick](<#func-pick>)
  - [PickBy](<#func-pickby>)
  - [Values](<#func-values>)

- **Concurrency and time related utility functions**
  - [After](<#func-after>)
  - [Before](<#func-before>)
  - [Delay](<#func-delay>)
  - [Flip](<#func-flip>)
  - [Memoize](<#func-memoizert-v-memoize>)
  - [NewDebounce](<#func-newdebounce>)
  - [Once](<#func-once>)
  - [Retry](<#func-rtypet-retry>)
  - [RetryWithDelay](<#func-rtypet-retrywithdelay>)

## func Abs

## func [Abs](<https://github.com/esimov/gogu/blob/master/math.go#L30>)

```go
func Abs[T Number](x T) T
```

Abs returns the absolut value of x.

## func [After](<https://github.com/esimov/gogu/blob/master/func.go#L29>)

```go
func After[V constraints.Signed](n *V, fn func())
```

After creates a function wrapper that does nothing at first. From the nth call onwards, it starts actually invoking the callback function. Useful for grouping responses, where you need to be sure that all the calls have finished just before proceeding to the actual job.

<details><summary>Example</summary>
<p>

```go
{
	sample := []int{1, 2, 3, 4, 5, 6}
	length := len(sample) - 1

	initVal := 0
	fn := func(val int) int {
		return val + 1
	}

	ForEach(sample, func(val int) {
		now := time.Now()
		After(&length, func() {
			<-time.After(10 * time.Millisecond)
			initVal = fn(initVal)
			after := time.Since(now).Milliseconds()
			fmt.Println(after)
		})
	})

}
```

#### Output

```
10
```
</p>
</details>

## func [Before](<https://github.com/esimov/gogu/blob/master/func.go#L39>)

```go
func Before[S ~string, T any, V constraints.Signed](n *V, c *cache.Cache[S, T], fn func() T) T
```

Before creates a function wrapper that memoizes its return value. From the nth call onwards, the memoized result of the last invocation is returned immediately instead of invoking function again. So the wrapper will invoke function at most n\-1 times.

<details><summary>Example</summary>
<p>

```go
{
  c := cache.New[string, int](cache.DefaultExpiration, cache.NoExpiration)

	var n = 3
	sample := []int{1, 2, 3}
	ForEach(sample, func(val int) {
		fn := func() int {
			<-time.After(10 * time.Millisecond)
			return n
		}
		res := Before(&n, c, fn)
		// The trick to test this function is to decrease the n value after each iteration.
		// We can be sure that the callback function is not served from the cache if n > 0.
		// In this case the cache item "func" should be empty.
		if n > 0 {
			val, _ := c.Get("func")
			fmt.Println(val)
			fmt.Println(res)
		}
		if n <= 0 {
			// Here the callback function is served from the cache.
			val, _ := c.Get("func")
			fmt.Println(val)
			fmt.Println(res)
		}
	})
}
```

#### Output

```
<nil>
2
<nil>
1
&{0 0}
0
```

</p>
</details>

## func [CamelCase](<https://github.com/esimov/gogu/blob/master/string.go#L96>)

```go
func CamelCase[T ~string](str T) T
```

CamelCase converts a string to camelCase \(https://en.wikipedia.org/wiki/CamelCase\).

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(CamelCase("Foo Bar"))
	fmt.Println(CamelCase("--foo-Bar--"))
	fmt.Println(CamelCase("__foo-_Bar__"))
	fmt.Println(CamelCase("__FOO BAR__"))
	fmt.Println(CamelCase(" FOO BAR "))
	fmt.Println(CamelCase("&FOO&baR "))
	fmt.Println(CamelCase("&&foo&&bar__"))
}
```

#### Output

```
fooBar
fooBar
fooBar
fooBar
fooBar
fooBar
fooBar
```

</p>
</details>

## func [Capitalize](<https://github.com/esimov/gogu/blob/master/string.go#L81>)

```go
func Capitalize[T ~string](str T) T
```

Capitalize converts the first letter of the string to uppercase and the remaining letters to lowercase.

## func [Chunk](<https://github.com/esimov/gogu/blob/master/slice.go#L405>)

```go
func Chunk[T comparable](slice []T, size int) [][]T
```

Chunk split the slice into groups of slices each having the length of size. In case the source slice cannot be distributed equally, the last slice will contain fewer elements.

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(Chunk([]int{0, 1, 2, 3}, 2))
	fmt.Println(Chunk([]int{0, 1, 2, 3, 4}, 2))
	fmt.Println(Chunk([]int{0, 1}, 1))

}
```

#### Output

```
[[0 1] [2 3]]
[[0 1] [2 3] [4]]
[[0] [1]]
```

</p>
</details>

## func [Clamp](<https://github.com/esimov/gogu/blob/master/math.go#L38>)

```go
func Clamp[T Number](num, min, max T) T
```

Clamp returns a range\-limited number between min and max.

## func [Compare](<https://github.com/esimov/gogu/blob/master/generic.go#L11>)

```go
func Compare[T comparable](a, b T, comp CompFn[T]) int
```

Compare compares two values using as comparator the callback function argument.

<details><summary>Example</summary>
<p>

```go
{
	res1 := Compare(1, 2, func(a, b int) bool {
		return a < b
	})
	fmt.Println(res1)

	res2 := Compare("a", "b", func(a, b string) bool {
		return a > b
	})
	fmt.Println(res2)

}
```

#### Output

```
1
-1
```

</p>
</details>

## func [Contains](<https://github.com/esimov/gogu/blob/master/slice.go#L172>)

```go
func Contains[T comparable](slice []T, value T) bool
```

Contains returns true if the value is present in the collection.

## func [Delay](<https://github.com/esimov/gogu/blob/master/func.go#L20>)

```go
func Delay(delay time.Duration, fn func()) *time.Timer
```

Delay invokes the callback function with a predefined delay.

<details><summary>Example</summary>
<p>

```go
{
	ch := make(chan struct{})
	now := time.Now()

	var value uint32
	timer := Delay(20*time.Millisecond, func() {
		atomic.AddUint32(&value, 1)
		ch <- struct{}{}
	})
	r1 := atomic.LoadUint32(&value)
	fmt.Println(r1)
	<-ch
	if timer.Stop() {
		<-timer.C
	}
	r1 = atomic.LoadUint32(&value)
	fmt.Println(r1)
	after := time.Since(now).Milliseconds()
	fmt.Println(after)

}
```

#### Output

```
0
1
20
```

</p>
</details>

## func [Difference](<https://github.com/esimov/gogu/blob/master/slice.go#L363>)

```go
func Difference[T comparable](s1, s2 []T) []T
```

Difference is similar to Without, but returns the values from the first slice that are not present in the second slice.

## func [DifferenceBy](<https://github.com/esimov/gogu/blob/master/slice.go#L384>)

```go
func DifferenceBy[T comparable](s1, s2 []T, fn func(T) T) []T
```

DifferenceBy is like Difference, except that invokes a callback function on each element of the slice, applying the criteria by which the difference is computed.

## func [Drop](<https://github.com/esimov/gogu/blob/master/slice.go#L425>)

```go
func Drop[T any](slice []T, n int) []T
```

Drop creates a new slice with n elements dropped from the beginning. If n \< 0 the elements will be dropped from the back of the collection.

## func [DropWhile](<https://github.com/esimov/gogu/blob/master/slice.go#L438>)

```go
func DropWhile[T any](slice []T, fn func(T) bool) []T
```

DropWhile creates a new slice excluding the elements dropped from the beginning. Elements are dropped by applying the condition invoked in the callback function.

<details><summary>Example</summary>
<p>

```go
{
	res := DropWhile([]string{"a", "aa", "bbb", "ccc"}, func(elem string) bool {
		return len(elem) > 2
	})
	fmt.Println(res)

}
```

#### Output

```
[a aa]
```

</p>
</details>

## func [DropRightWhile](<https://github.com/esimov/gogu/blob/master/slice.go#L452>)

```go
func DropRightWhile[T any](slice []T, fn func(T) bool) []T
```

DropRightWhile creates a new slice excluding the elements dropped from the end. Elements are dropped by applying the condition invoked in the callback function.

## func [Duplicate](<https://github.com/esimov/gogu/blob/master/slice.go#L182>)

```go
func Duplicate[T comparable](slice []T) []T
```

Duplicate returns the duplicated values of a collection.

<details><summary>Example</summary>
<p>

```go
{
	input := []int{-1, -1, 0, 1, 2, 3, 2, 5, 1, 6}
	fmt.Println(Duplicate(input))
}
```

#### Output

```
[-1 1 2]
```

</p>
</details>

## func [DuplicateWithIndex](<https://github.com/esimov/gogu/blob/master/slice.go#L206>)

```go
func DuplicateWithIndex[T comparable](slice []T) map[T]int
```

DuplicateWithIndex puts the duplicated values of a collection into a map as a key value pair, where the key is the collection element and the value is its position.

<details><summary>Example</summary>
<p>

```go
{
	input := []int{-1, -1, 0, 1, 2, 3, 2, 5, 1, 6}
	fmt.Println(DuplicateWithIndex(input))

}
```

#### Output

```
map[-1:0 1:3 2:4]
```

</p>
</details>

## func [Equal](<https://github.com/esimov/gogu/blob/master/generic.go#L21>)

```go
func Equal[T comparable](a, b T) bool
```

Equal checks if two values are equal.

## func [Every](<https://github.com/esimov/gogu/blob/master/slice.go#L136>)

```go
func Every[T any](slice []T, fn func(T) bool) bool
```

Every returns true if all the elements of a slice satisfies the criteria of the callback function.

## func [Filter](<https://github.com/esimov/gogu/blob/master/filter.go#L4>)

```go
func Filter[T any](slice []T, fn func(T) bool) []T
```

Filter returns all the elements from the collection which satisfies the conditional logic of the callback function.

<details><summary>Example</summary>
<p>

```go
{
	input := []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50}
	res := Filter(input, func(val int) bool {
		return val >= 10
	})
	fmt.Println(res)
}
```

#### Output

```
[10 20 30 40 50]
```

</p>
</details>

## func [Filter2DMapCollection](<https://github.com/esimov/gogu/blob/master/filter.go#L63>)

```go
func Filter2DMapCollection[K comparable, V any](collection []map[K]map[K]V, fn func(map[K]V) bool) []map[K]map[K]V
```

Filter2DMapCollection filter out a two\-dimensional collection of map items by applying the conditional logic of the callback function.

## func [FilterMap](<https://github.com/esimov/gogu/blob/master/filter.go#L33>)

```go
func FilterMap[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V
```

FilterMap iterates over the elements of a collection and returns a new collection representing all the items which satisfies the criteria formulated in the callback function.

<details><summary>Example</summary>
<p>

```go
{
	input := map[int]string{1: "John", 2: "Doe", 3: "Fred"}
	res := FilterMap(input, func(v string) bool {
		return v == "John"
	})
	fmt.Println(res)
}
```

#### Output

```
map[1:John]
```

</p>
</details>

## func [FilterMapCollection](<https://github.com/esimov/gogu/blob/master/filter.go#L47>)

```go
func FilterMapCollection[K comparable, V any](collection []map[K]V, fn func(V) bool) []map[K]V
```

FilterMapCollection filter out a one dimensional collection of map items by applying the conditional logic of the callback function.

<details><summary>Example</summary>
<p>

```go
{
	input := []map[string]int{
		{"bernie": 22},
		{"robert": 30},
	}
	res := FilterMapCollection(input, func(val int) bool {
		return val > 22
	})
	fmt.Println(res)

}
```

#### Output

```
[map[robert:30]]
```

</p>
</details>

## func [Find](<https://github.com/esimov/gogu/blob/master/map.go#L128>)

```go
func Find[K constraints.Ordered, V any](m map[K]V, fn func(V) bool) map[K]V
```

Find iterates over the elements of a map and returns the first item for which the callback function returns true.

## func [FindAll](<https://github.com/esimov/gogu/blob/master/find.go#L33>)

```go
func FindAll[T any](s []T, fn func(T) bool) map[int]T
```

FindAll is like FindIndex, but returns into a map all the values which satisfies the conditional logic of the callback function. The map key represents the position of the found value and the value is the item itself.

<details><summary>Example</summary>
<p>

```go
{
	input := []int{1, 2, 3, 4, 2, -2, -1, 2}
	items := FindAll(input, func(v int) bool {
		return v == 2
	})
	fmt.Println(items)

}
```

#### Output

```
map[1:2 4:2 7:2]
```

</p>
</details>

## func [FindByKey](<https://github.com/esimov/gogu/blob/master/map.go#L167>)

```go
func FindByKey[K comparable, V any](m map[K]V, fn func(K) bool) map[K]V
```

FindByKey is like Find, but returns the first item for which the callback function returns true.

## func [FindIndex](<https://github.com/esimov/gogu/blob/master/find.go#L11>)

```go
func FindIndex[T any](s []T, fn func(T) bool) int
```

FindIndex returns the index of the first found element.

## func [FindKey](<https://github.com/esimov/gogu/blob/master/map.go#L155>)

```go
func FindKey[K comparable, V any](m map[K]V, fn func(V) bool) K
```

FindKey is like Find, but returns the first item key position for which the callback function returns true.

## func [FindLastIndex](<https://github.com/esimov/gogu/blob/master/find.go#L21>)

```go
func FindLastIndex[T any](s []T, fn func(T) bool) int
```

FindLastIndex is like FindIndex, only that returns the index of last found element.

## func [FindMax](<https://github.com/esimov/gogu/blob/master/find.go#L103>)

```go
func FindMax[T constraints.Ordered](s []T) T
```

FindMax finds the maximum value of a slice.

## func [FindMaxBy](<https://github.com/esimov/gogu/blob/master/find.go#L120>)

```go
func FindMaxBy[T constraints.Ordered](s []T, fn func(val T) T) T
```

FindMaxBy is like FindMax except that it accept a callback function and the conditional logic is applied over the resulted value. If there are more than one identical values resulted from the callback function the first one is returned.

## func [FindMaxByKey](<https://github.com/esimov/gogu/blob/master/find.go#L135>)

```go
func FindMaxByKey[K comparable, T constraints.Ordered](mapSlice []map[K]T, key K) (T, error)
```

FindMaxByKey finds the maximum value from a map by using some existing key as a parameter.

## func [FindMin](<https://github.com/esimov/gogu/blob/master/find.go#L46>)

```go
func FindMin[T constraints.Ordered](s []T) T
```

FindMin finds the minimum value of a slice.

## func [FindMinBy](<https://github.com/esimov/gogu/blob/master/find.go#L63>)

```go
func FindMinBy[T constraints.Ordered](s []T, fn func(val T) T) T
```

FindMinBy is like FindMin except that it accept a callback function and the conditional logic is applied over the resulted value. If there are more than one identical values resulted from the callback function the first one is used.

## func [FindMinByKey](<https://github.com/esimov/gogu/blob/master/find.go#L78>)

```go
func FindMinByKey[K comparable, T constraints.Ordered](mapSlice []map[K]T, key K) (T, error)
```

FindMinByKey finds the minimum value from a map by using some existing key as a parameter.

## func [Flatten](<https://github.com/esimov/gogu/blob/master/slice.go#L248>)

```go
func Flatten[T any](slice any) ([]T, error)
```

Flatten flattens the slice all the way down to the deepest nesting level.

<details><summary>Example</summary>
<p>

```go
{
	input := []any{[]int{1, 2, 3}, []any{[]int{4}, 5}}
	result, _ := Flatten[int](input)
	fmt.Println(result)

}
```

#### Output

```
[1 2 3 4 5]
```

</p>
</details>

## func [Flip](<https://github.com/esimov/gogu/blob/master/func.go#L13>)

```go
func Flip[T any](fn func(args ...T) []T) func(args ...T) []T
```

Flip creates a function that invokes fn with arguments reversed.

<details><summary>Example</summary>
<p>

```go
{
	flipped := Flip(func(args ...int) []int {
		return ToSlice(args...)
	})
	fmt.Println(flipped(1, 2, 3))

}
```

#### Output

```
[3 2 1]
```

</p>
</details>

## func [ForEach](<https://github.com/esimov/gogu/blob/master/slice.go#L69>)

```go
func ForEach[T any](slice []T, fn func(T))
```

ForEach iterates over the elements of a collection and invokes the callback fn function on each element.

<details><summary>Example</summary>
<p>

```go
{
	input := []int{1, 2, 3, 4}
	output := []int{}

	ForEach(input, func(val int) {
		val = val * 2
		output = append(output, val)
	})
	fmt.Println(output)

}
```

#### Output

```
[2 4 6 8]
```

</p>
</details>

## func [ForEachRight](<https://github.com/esimov/gogu/blob/master/slice.go#L76>)

```go
func ForEachRight[T any](slice []T, fn func(T))
```

ForEachRight is the same as ForEach, but starts the iteration from the last element.

## func [GroupBy](<https://github.com/esimov/gogu/blob/master/slice.go#L481>)

```go
func GroupBy[T1, T2 comparable](slice []T1, fn func(T1) T2) map[T2][]T1
```

GroupBy splits a collection into a key\-value set, grouped by the result of running each value through the callback function fn. The return value is a map where the key is the conditional logic of the callback function and the values are the callback function returned values.

<details><summary>Example</summary>
<p>

```go
{
	input := []float64{1.3, 1.5, 2.1, 2.9}
	res := GroupBy(input, func(val float64) float64 {
		return math.Floor(val)
	})
	fmt.Println(res)

}
```

#### Output

```
map[1:[1.3 1.5] 2:[2.1 2.9]]
```

</p>
</details>

## func [InRange](<https://github.com/esimov/gogu/blob/master/math.go#L48>)

```go
func InRange[T Number](num, lo, up T) bool
```

InRange checks if a number is inside a range.

## func [IndexOf](<https://github.com/esimov/gogu/blob/master/slice.go#L38>)

```go
func IndexOf[T comparable](s []T, val T) int
```

IndexOf returns the index of the firs occurrence of a value in the slice, or \-1 if value is not present in the slice.

## func [Intersection](<https://github.com/esimov/gogu/blob/master/slice.go#L286>)

```go
func Intersection[T comparable](params ...[]T) []T
```

Intersection computes the list of values that are the intersection of all the slices. Each value in the result should be present in each of the provided slices.

<details><summary>Example</summary>
<p>

```go
{
	res1 := Intersection([]int{1, 2, 4}, []int{0, 2, 1}, []int{2, 1, -2})
	fmt.Println(res1)

	res2 := Intersection([]string{"a", "b"}, []string{"a", "a", "a"}, []string{"b", "a", "e"})
	fmt.Println(res2)

}
```

#### Output

```
[1 2]
[a]
```

</p>
</details>

## func [IntersectionBy](<https://github.com/esimov/gogu/blob/master/slice.go#L310>)

```go
func IntersectionBy[T comparable](fn func(T) T, params ...[]T) []T
```

IntersectionBy is like Intersection, except that it accepts and callback function which is invoked on each element of the collection.

<details><summary>Example</summary>
<p>

```go
{
	result1 := IntersectionBy(func(v float64) float64 {
		return math.Floor(v)
	}, []float64{2.1, 1.2}, []float64{2.3, 3.4}, []float64{1.0, 2.3})
	fmt.Println(result1)

	result2 := IntersectionBy(func(v int) int {
		return v % 2
	}, []int{1, 2}, []int{2, 1})
	fmt.Println(result2)

}
```

#### Output

```
[2.1]
[1 2]
```

</p>
</details>

## func [Invert](<https://github.com/esimov/gogu/blob/master/map.go#L180>)

```go
func Invert[K, V comparable](m map[K]V) map[V]K
```

Invert returns a copy of the map where the keys become the values and the values the keys. For this to work, all of your map's values should be unique.

## func [KebabCase](<https://github.com/esimov/gogu/blob/master/string.go#L132>)

```go
func KebabCase[T ~string](str T) T
```

KebabCase converts a string to kebab\-case \(https://en.wikipedia.org/wiki/Letter_case#Kebab_case\).

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(KebabCase("fooBarBaz"))
	fmt.Println(KebabCase("Foo BarBaz"))
	fmt.Println(KebabCase("Foo_Bar_Baz"))

}
```

#### Output

```
foo-bar-baz
foo-bar-baz
foo-bar-baz
```

</p>
</details>

## func Keys

```go
func Keys[K comparable, V any](m map[K]V) []K
```

Keys retrieve all the existing keys of a map.

## func [LastIndexOf](<https://github.com/esimov/gogu/blob/master/slice.go#L49>)

```go
func LastIndexOf[T comparable](s []T, val T) int
```

LastIndexOf returns the index of the last occurrence of a value.

## func [Less](<https://github.com/esimov/gogu/blob/master/generic.go#L26>)

```go
func Less[T constraints.Ordered](a, b T) bool
```

Less checks if the first value is less than the second.

## func [Map](<https://github.com/esimov/gogu/blob/master/slice.go#L59>)

```go
func Map[T1, T2 any](slice []T1, fn func(T1) T2) []T2
```

Map produces a new slice of values by mapping each value in the list through a transformation function.

<details><summary>Example</summary>
<p>

```go
{
	res := Map([]int{1, 2, 3}, func(val int) int {
		return val * 2
	})
	fmt.Println()

}
```

#### Output

```
[2 4 6]
```
</p>
</details>

## func [MapCollection](<https://github.com/esimov/gogu/blob/master/map.go#L116>)

```go
func MapCollection[K comparable, V any](m map[K]V, fn func(V) V) []V
```

MapCollection is like the Map method, but applied to maps. It runs each element of the map over an iteratee function and saves the resulted values into a new map.

## func [MapContains](<https://github.com/esimov/gogu/blob/master/map.go#L89>)

```go
func MapContains[K, V comparable](m map[K]V, value V) bool
```

MapContains returns true if the value is present in the list otherwise false.

## func [MapEvery](<https://github.com/esimov/gogu/blob/master/map.go#L67>)

```go
func MapEvery[K comparable, V any](m map[K]V, fn func(V) bool) bool
```

MapEvery returns true if all the elements of a map satisfies the criteria of the callback function.

## func [MapKeys](<https://github.com/esimov/gogu/blob/master/map.go#L56>)

```go
func MapKeys[K comparable, V any, R comparable](m map[K]V, fn func(K, V) R) map[R]V
```

MapKeys is the opposite of MapValues. It creates a new map with the same number of elements as the original one, but this time the callback function \(fn\) is invoked over the map keys.

## func [MapSome](<https://github.com/esimov/gogu/blob/master/map.go#L78>)

```go
func MapSome[K comparable, V any](m map[K]V, fn func(V) bool) bool
```

MapSome returns true if some elements of a map satisfies the criteria of the callback function.

## func [MapUnique](<https://github.com/esimov/gogu/blob/master/map.go#L99>)

```go
func MapUnique[K, V comparable](m map[K]V) map[K]V
```

MapUnique removes the duplicate values from a map.

## func [MapValues](<https://github.com/esimov/gogu/blob/master/map.go#L44>)

```go
func MapValues[K comparable, V, R any](m map[K]V, fn func(V) R) map[K]R
```

MapValues creates a new map with the same number of elements as the original one, but running each map value through a callback function \(fn\).

## func [Max](<https://github.com/esimov/gogu/blob/master/math.go#L18>)

```go
func Max[T constraints.Ordered](values ...T) T
```

Max returns the biggest value from the provided parameters.

## func [Mean](<https://github.com/esimov/gogu/blob/master/slice.go#L28>)

```go
func Mean[T Number](slice []T) T
```

Mean computes the mean value of the slice elements.

## func [Merge](<https://github.com/esimov/gogu/blob/master/slice.go#L236>)

```go
func Merge[T any](s []T, params ...[]T) []T
```

Merge merges the first slice with the other slices defined as variadic parameter.

## func [Min](<https://github.com/esimov/gogu/blob/master/math.go#L6>)

```go
func Min[T constraints.Ordered](values ...T) T
```

Min returns the lowest value from the provided parameters.

## func [N](<https://github.com/esimov/gogu/blob/master/range.go#L79>)

```go
func N[T Number](s string) (T, error)
```

N converts a string to a generic number.

## func [NewDebounce](<https://github.com/esimov/gogu/blob/master/func.go#L125>)

```go
func NewDebounce(wait time.Duration) (func(f func()), func())
```

NewDebounce creates a new debounced version of the invoked function which postpone the execution with a time delay passed in as a function argument. It returns a callback function which will be invoked after the predefined delay and also a cancel method which should be invoked to cancel a scheduled debounce.

<details><summary>Example</summary>
<p>

```go
{
	var (
		counter1 uint64
		counter2 uint64
	)

	f1 := func() {
		atomic.AddUint64(&counter1, 1)
	}

	f2 := func() {
		atomic.AddUint64(&counter2, 1)
	}

	debounce, cancel := NewDebounce(10 * time.Millisecond)
	for i := 0; i < 2; i++ {
		for j := 0; j < 100; j++ {
			debounce(f1)
		}
		<-time.After(20 * time.Millisecond)
	}
	cancel()

	debounce, cancel = NewDebounce(10 * time.Millisecond)
	for i := 0; i < 5; i++ {
		for j := 0; j < 50; j++ {
			debounce(f2)
		}
		for j := 0; j < 50; j++ {
			debounce(f2)
		}
		<-time.After(20 * time.Millisecond)
	}
	cancel()

	c1 := atomic.LoadUint64(&counter1)
	c2 := atomic.LoadUint64(&counter2)
	fmt.Println(c1)
	fmt.Println(c2)

}
```

#### Output

```
2
5
```

</p>
</details>

## func [Nth](<https://github.com/esimov/gogu/blob/master/find.go#L162>)

```go
func Nth[T any](slice []T, nth int) (T, error)
```

Nth returns the nth element of the collection. In case of negative value the nth element is returned from the end of the collection. In case nth is out of bounds an error is returned.

## func [Null](<https://github.com/esimov/gogu/blob/master/string.go#L10>)

```go
func Null[T any]() T
```

## func [NumToString](<https://github.com/esimov/gogu/blob/master/range.go#L108>)

```go
func NumToString[T Number](n T) string
```

NumToString converts a number to a string. In case of a number of type float \(float32|float64\) this will be rounded to 2 decimal places.

## func [Omit](<https://github.com/esimov/gogu/blob/master/map.go#L237>)

```go
func Omit[K comparable, V any](collection map[K]V, keys ...K) map[K]V
```

Omit is the opposite of Pick, it extracts all the map elements which keys are not omitted.

<details><summary>Example</summary>
<p>

```go
{
	res := Omit(map[string]any{"name": "moe", "age": 40, "active": false}, "name", "age")
	fmt.Println(res)

}
```

#### Output

```
map[active:false]
```

</p>
</details>

## func [OmitBy](<https://github.com/esimov/gogu/blob/master/map.go#L248>)

```go
func OmitBy[K comparable, V any](collection map[K]V, fn func(key K, val V) bool) map[K]V
```

OmitBy is the opposite of PickBy, it removes all the map elements for which the callback function returns true.

<details><summary>Example</summary>
<p>

```go
{
	res := OmitBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(key string, val int) bool {
		return val%2 == 1
	})
	fmt.Println(res)

}
```

#### Output

```
map[b:2]
```

</p>
</details>

## func [Once](<https://github.com/esimov/gogu/blob/master/func.go#L56>)

```go
func Once[S ~string, T comparable, V constraints.Signed](c *cache.Cache[S, T], fn func() T) T
```

Once is like Before, but it's invoked only once. Repeated calls to the modified function will have no effect and the function invocation is returned from the cache.

<details><summary>Example</summary>
<p>

```go
{
	c := cache.New[string, int](cache.DefaultExpiration, cache.NoExpiration)

	ForEach([]int{1, 2, 3, 4, 5}, func(val int) {
		fn := func(val int) func() int {
			<-time.After(10 * time.Millisecond)
			return func() int {
				return val
			}
		}
		res := Once[string, int, int](c, fn(val))

        // We can test the implementation correctness by invoking the `Once` function multiple times.
	    // When it's invoked for the first time the result should be served from the callback function.
	    // From the second invocation onward the results are served from the cache.
	    // In our example the results of each invokation should be always equal with 1.
		fmt.Println(res)
	})
	c.Flush()
}
```

#### Output

```
1
1
1
1
1
```

</p>
</details>

## func [Pad](<https://github.com/esimov/gogu/blob/master/string.go#L236>)

```go
func Pad[T ~string](str T, size int, token string) T
```

Pad pads string on the left and right sides if it's shorter than length. Padding characters are truncated if they can't be evenly divided by length.

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(Pad("abc", 2, "."))
	fmt.Println(Pad("abc", 3, "."))
	fmt.Println(Pad("abc", 4, "."))
	fmt.Println(Pad("abc", 5, "."))
}
```

#### Output

```
abc
abc
abc.
.abc.
```

</p>
</details>

## func [PadLeft](<https://github.com/esimov/gogu/blob/master/string.go#L194>)

```go
func PadLeft[T ~string](str T, size int, token string) T
```

PadLeft pads string on the left side if it's shorter than length. Padding characters are truncated if they exceed length.

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(PadLeft("abc", 8, "..."))
	fmt.Println(PadLeft("abc", 4, "_"))
	fmt.Println(PadLeft("abc", 6, "_-"))

}
```

#### Output

```
.....abc
_abc
_-_abc
```

</p>
</details>

## func [PadRight](<https://github.com/esimov/gogu/blob/master/string.go#L215>)

```go
func PadRight[T ~string](str T, size int, token string) T
```

PadRight pads string on the right side if it's shorter than length. Padding characters are truncated if they exceed length.

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(PadRight("abc", 8, "..."))
	fmt.Println(PadRight("abc", 6, "........"))
}
```

#### Output

```
abc.....
abc...
```

</p>
</details>

## func [Partition](<https://github.com/esimov/gogu/blob/master/slice.go#L157>)

```go
func Partition[T comparable](slice []T, fn func(T) bool) [2][]T
```

Partition splits the collection elements into two, the ones which satisfies the condition expressed in the callback function (`fn`) and those which does not satisfy the condition.

<details><summary>Example</summary>
<p>

```go
{
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res1 := Partition(input, func(val int) bool {
		return val >= 5
	})
	fmt.Println(res1)

	res2 := Partition(input, func(val int) bool {
		return val < 5
	})
	fmt.Println(res2)

}
```

#### Output

```
[[5 6 7 8 9] [0 1 2 3 4]]
[[0 1 2 3 4] [5 6 7 8 9]]
```

</p>
</details>

## func [PartitionMap](<https://github.com/esimov/gogu/blob/master/map.go#L260>)

```go
func PartitionMap[K comparable, V any](mapSlice []map[K]V, fn func(map[K]V) bool) [2][]map[K]V
```

PartitionMap split the collection into two arrays, the one whose elements satisfy the condition expressed in the callback function (`fn`) and one whose elements don't satisfy the condition.

## func [Pick](<https://github.com/esimov/gogu/blob/master/map.go#L208>)

```go
func Pick[K comparable, V any](collection map[K]V, keys ...K) (map[K]V, error)
```

Pick extracts the elements from the map which have the key defined in the allowed keys.

<details><summary>Example</summary>
<p>

```go
{
	res, _ := Pick(map[string]any{"name": "moe", "age": 20, "active": true}, "name", "age")
	fmt.Println(res)
}
```

#### Output

```
map[age:20 name:moe]
```

</p>
</details>

## func [PickBy](<https://github.com/esimov/gogu/blob/master/map.go#L224>)

```go
func PickBy[K comparable, V any](collection map[K]V, fn func(key K, val V) bool) map[K]V
```

PickBy extracts all the map elements for which the callback function returns truthy.

<details><summary>Example</summary>
<p>

```go
{
	res := PickBy(map[string]int{"aa": 1, "b": 2, "c": 3}, func(key string, val int) bool {
		return len(key) == 1
	})
	fmt.Println(res)

}
```

#### Output

```
map[b:2 c:3]
```

</p>
</details>

## func [Pluck](<https://github.com/esimov/gogu/blob/master/map.go#L192>)

```go
func Pluck[K comparable, V any](mapSlice []map[K]V, key K) []V
```

Pluck extracts all the values of a map by the key definition.

<details><summary>Example</summary>
<p>

```go
{
	input := []map[string]string{
		{"name": "moe", "email": "moe@example.com"},
		{"name": "larry", "email": "larry@example.com"},
		{"name": "curly", "email": "curly@example.com"},
		{"name": "moly", "email": "moly@example.com"},
	}
	res := Pluck(input, "name")
	fmt.Println(res)

}
```

#### Output

```
[moe larry curly moly]
```

</p>
</details>

## func [Range](<https://github.com/esimov/gogu/blob/master/range.go#L21>)

```go
func Range[T Number](args ...T) ([]T, error)
```

Range creates a slice of integers progressing from start up to, but not including end. This method can accept 1, 2 or 3 arguments. Depending on the number of provided parameters, \`start\`, \`step\` and \`end\` has the following meaning:

\[start=0\]: The start of the range. If omitted it defaults to 0.

\[step=1\]: The value to increment or decrement by.

end: The end of the range.

In case you'd like negative values, use a negative step.

<details><summary>Example</summary>
<p>

```go
{
	r1, _ := Range(5)
	r2, _ := Range(1, 5)
	r3, _ := Range(0, 2, 10)
	r4, _ := Range(-4)
	r5, _ := Range(-1, -4)
	r6, _ := Range(0, -1, -4)
	r7, _ := Range[float64](0, 0.12, 0.9)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(r5)
	fmt.Println(r6)
	fmt.Println(r7)

}
```

#### Output

```
[0 1 2 3 4]
[1 2 3 4]
[0 2 4 6 8]
[0 -1 -2 -3]
[-1 -2 -3]
[0 -1 -2 -3]
[0 0.12 0.24 0.36 0.48 0.6 0.72 0.84]
```

</p>
</details>

## func [RangeRight](<https://github.com/esimov/gogu/blob/master/range.go#L70>)

```go
func RangeRight[T Number](params ...T) ([]T, error)
```

RangeRight is like Range, only that it populates the slice in descending order.

## func [Reduce](<https://github.com/esimov/gogu/blob/master/slice.go#L84>)

```go
func Reduce[T1, T2 any](slice []T1, fn func(T1, T2) T2, initVal T2) T2
```

Reduce reduces the collection to a value which is the accumulated result of running each element in the collection through the callback function yielding a single value.

<details><summary>Example</summary>
<p>

```go
{
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

}
```

#### Output

```
10
abcd
```

</p>
</details>

## func [Reject](<https://github.com/esimov/gogu/blob/master/filter.go#L18>)

```go
func Reject[T any](slice []T, fn func(val T) bool) []T
```

Reject is the opposite of Filter. It returns the values from the collection without the elements for which the callback function returns true.

<details><summary>Example</summary>
<p>

```go
{
	input := []int{1, 2, 3, 4, 5, 6, 10, 20, 30, 40, 50}
	res = Reject(input, func(val int) bool {
		return val >= 10
	})
	fmt.Println(res)
}
```

#### Output

```
[1 2 3 4 5 6]
```

</p>
</details>

## func [Reverse](<https://github.com/esimov/gogu/blob/master/slice.go#L96>)

```go
func Reverse[T any](sl []T) []T
```

Reverse reverses the order of elements, so that the first element becomes the last, the second element becomes the second to last, and so on.

## func [ReverseStr](<https://github.com/esimov/gogu/blob/master/string.go#L322>)

```go
func ReverseStr[T ~string](str T) T
```

ReverseStr returns a new string with the characters in reverse order.

## func [Shuffle](<https://github.com/esimov/gogu/blob/master/shuffle.go#L8>)

```go
func Shuffle[T any](src []T) []T
```

Shuffle implements the Fisher\-Yates shuffle algorithm applied to a slice.

## func [SliceToMap](<https://github.com/esimov/gogu/blob/master/map.go#L281>)

```go
func SliceToMap[K comparable, T any](s1 []K, s2 []T) map[K]T
```

SliceToMap converts a slice to a map. It panics in case the parameter slices length are not identical. The map keys will be the items from the first slice and the values the items from the second slice.

## func [SnakeCase](<https://github.com/esimov/gogu/blob/master/string.go#L127>)

```go
func SnakeCase[T ~string](str T) T
```

SnakeCase converts a string to snake\_case \(https://en.wikipedia.org/wiki/Snake_case\).

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(SnakeCase("fooBarBaz"))
	fmt.Println(SnakeCase("Foo BarBaz"))
	fmt.Println(SnakeCase("Foo_Bar_Baz"))

}
```

#### Output

```
foo_bar_baz
foo_bar_baz
foo_bar_baz
```

</p>
</details>

## func [Some](<https://github.com/esimov/gogu/blob/master/slice.go#L146>)

```go
func Some[T any](slice []T, fn func(T) bool) bool
```

Some returns true if some elements of a slice satisfies the criteria of the callback function.

## func [SplitAtIndex](<https://github.com/esimov/gogu/blob/master/string.go#L265>)

```go
func SplitAtIndex[T ~string](str T, index int) []T
```

SplitAtIndex split the string at the specified index and returns a slice with the resulted two substrings.

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(SplitAtIndex("abcdef", -1))
	fmt.Println(SplitAtIndex("abcdef", 0))
	fmt.Println(SplitAtIndex("abcdef", 1))
	fmt.Println(SplitAtIndex("abcdef", 2))
	fmt.Println(SplitAtIndex("abcdef", 5))
	fmt.Println(SplitAtIndex("abcdef", 6))

}
```

#### Output

```
[ abcdef]
[a bcdef]
[ab cdef]
[abc def]
[abcdef ]
[abcdef ]
```

</p>
</details>

## func [Substr](<https://github.com/esimov/gogu/blob/master/string.go#L27>)

```go
func Substr[T ~string](str T, offset, length int) T
```

Substr returns the portion of string specified by the offset and length.

If offset is non\-negative, the returned string will start at the offset'th position in string, counting from zero.

If offset is negative, the returned string will start at the offset'th character from the end of string.

If string is less than offset characters long, an empty string will be returned.

If length is negative, then that many characters will be omitted from the end of string starting from the offset position.

<details><summary>Example</summary>
<p>

```go
{
	str1 := Substr("abcdef", 0, 0)
	str2 := Substr("abcdef", -1, 0)
	str3 := Substr("abcdef", 7, 7)
	str4 := Substr("abcdef", 0, 20)
	str5 := Substr("abcdef", 5, 10)
	str6 := Substr("abcdef", 0, -1)
	str7 := Substr("abcdef", 2, -1)
	str8 := Substr("abcdef", 4, -4)
	str9 := Substr("abcdef", -3, -1)
	str10 := Substr("abcdef", 1, 3)

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)
	fmt.Println(str7)
	fmt.Println(str8)
	fmt.Println(str9)
	fmt.Println(str10)

}
```

#### Output

```
abcdef
f
abcde
cde

de
bcd
```

</p>
</details>

## func [Sum](<https://github.com/esimov/gogu/blob/master/slice.go#L9>)

```go
func Sum[T Number](slice []T) T
```

Sum returns the sum of the slice items. These have to satisfy the type constraints declared as Number.

## func [SumBy](<https://github.com/esimov/gogu/blob/master/slice.go#L19>)

```go
func SumBy[T1 any, T2 Number](slice []T1, fn func(T1) T2) T2
```

SumBy is like Sum except it accept a callback function which is invoked for each element in the slice to generate the value to be summed.

## func [ToLower](<https://github.com/esimov/gogu/blob/master/string.go#L58>)

```go
func ToLower[T ~string](str T) T
```

ToLower converts a string to Lowercase.

## func [ToSlice](<https://github.com/esimov/gogu/blob/master/slice.go#L544>)

```go
func ToSlice[T any](args ...T) []T
```

ToSlice returns the function arguments as a slice.

## func [ToUpper](<https://github.com/esimov/gogu/blob/master/string.go#L69>)

```go
func ToUpper[T ~string](str T) T
```

ToUpper converts a string to Uppercase.

## func Union

```go
func Union[T comparable](slice any) ([]T, error)
```

Union computes the union of the passed\\\-in slice and returns an ordered list of unique items that are present in one or more of the slices.

<details><summary>Example</summary>
<p>

```go
{
	input := []any{[]any{1, 2, []any{3, []int{4, 5, 6}}}, 7, []int{1, 2}, 3, []int{4, 7}, 8, 9, 9}
	res, _ := Union[int](input)
	fmt.Println(res)

}
```

#### Output

```
[1 2 3 4 5 6 7 8 9]
```

</p>
</details>

## func [Unique](<https://github.com/esimov/gogu/blob/master/slice.go#L105>)

```go
func Unique[T comparable](slice []T) []T
```

Unique returns the collection unique values.

## func [UniqueBy](<https://github.com/esimov/gogu/blob/master/slice.go#L121>)

```go
func UniqueBy[T comparable](slice []T, fn func(T) T) []T
```

UniqueBy is like Unique except that it accept a callback function which is invoked on each element of the slice applying the criteria by which the uniqueness is computed.

## func [Unwrap](<https://github.com/esimov/gogu/blob/master/string.go#L297>)

```go
func Unwrap[T ~string](str T, token string) T
```

Unwrap a string with the specified token.

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(Unwrap("'abc'", "'"))
	fmt.Println(Unwrap("*abc*", "*"))
	fmt.Println(Unwrap("*a*bc*", "*"))
	fmt.Println(Unwrap("''abc''", "''"))
	fmt.Println(Unwrap("\"abc\"", "\""))

}
```

#### Output

```
abc
abc
a*bc
abc
abc
```

</p>
</details>

## func [Unzip](<https://github.com/esimov/gogu/blob/master/slice.go#L516>)

```go
func Unzip[T any](slices ...[]T) [][]T
```

Unzip is the opposite of Zip: given a slice of slices it returns a series of new slices, the first of which contains all the first elements in the input slices, the second of which contains all the second elements, and so on.

<details><summary>Example</summary>
<p>

```go
{
	res := Unzip([]any{"one", 1}, []any{"two", 2})
	fmt.Println(res)

}
```

#### Output

```
[[one two] [1 2]]
```

</p>
</details>

## func [Values](<https://github.com/esimov/gogu/blob/master/map.go#L30>)

```go
func Values[K comparable, V any](m map[K]V) []V
```

Values retrieve all the existing values of a map.

## func [Without](<https://github.com/esimov/gogu/blob/master/slice.go#L342>)

```go
func Without[T1 comparable, T2 any](slice []T1, values ...T1) []T1
```

Without returns a copy of the slice with all the values defined in the variadic parameter removed.

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(Without[int, int]([]int{2, 1, 2, 3}, 1, 2))
	fmt.Println(Without[int, int]([]int{1, 2, 3, 4}, 3, 4))
	fmt.Println(Without[int, int]([]int{0, 1, 2, 3, 4, 5}, 0, 3, 4, 5))
	fmt.Println(Without[float64, float64]([]float64{1.0, 2.2, 3.0, 4.2}, 3.0, 4.2))

}
```

#### Output

```
[3]
[1 2]
[1 2]
[1 2.2]
```

</p>
</details>

## func [Wrap](<https://github.com/esimov/gogu/blob/master/string.go#L286>)

```go
func Wrap[T ~string](str T, token string) T
```

Wrap a string with the specified token.

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(Unwrap("'abc'", "'"))
	fmt.Println(Unwrap("*abc*", "*"))
	fmt.Println(Unwrap("*a*bc*", "*"))
	fmt.Println(Unwrap("''abc''", "''"))
	fmt.Println(Unwrap("\"abc\"", "\""))

}
```

#### Output

```
abc
abc
a*bc
abc
abc
```

</p>
</details>

## func [WrapAllRune](<https://github.com/esimov/gogu/blob/master/string.go#L309>)

```go
func WrapAllRune[T ~string](str T, token string) T
```

WrapAllRune is like Wrap, only that it's applied over runes instead of strings.

<details><summary>Example</summary>
<p>

```go
{
	fmt.Println(WrapAllRune("abc", ""))
	fmt.Println(WrapAllRune("abc", "'"))
	fmt.Println(WrapAllRune("abc", "*"))
	fmt.Println(WrapAllRune("abc", "-"))

}
```

#### Output

```
abc
'a''b''c'
*a**b**c*
-a--b--c-
```

</p>
</details>

## func [Zip](<https://github.com/esimov/gogu/blob/master/slice.go#L486>)

```go
func Zip[T any](slices ...[]T) [][]T
```

Zip iteratively merges together the values of the slice parameters with the values at the corresponding position.

<details><summary>Example</summary>
<p>

```go
{
	res := Zip([]any{"one", "two"}, []any{1, 2})
	fmt.Println(res)

}
```

#### Output

```
[[one 1] [two 2]]
```

</p>
</details>

## func [ToMap](<https://github.com/esimov/gogu/blob/master/slice.go#L552>)

```go
func ToMap[T, K comparable](slice []T, fn func(x T) K) map[K]T
```

Converts slice to map, where keys are generated with `fn` function.

<details><summary>Example</summary>
<p>

```go
{
	slice := []int{1, 2, 3}
	result := ToMap(slice, func(x int) int { return x })
}
```

#### Output

```
{1: 1, 2: 2, 3: 3}
```

</p>
</details>





## type [Bound](<https://github.com/esimov/gogu/blob/master/find.go#L180-L182>)

```go
type Bound[T constraints.Signed] struct {
    Min, Max T
}
```

### func \(Bound\[T\]\) [Enclose](<https://github.com/esimov/gogu/blob/master/find.go#L185>)

```go
func (b Bound[T]) Enclose(nth T) bool
```

Enclose checks if an element is inside the bounds.

## type [CompFn](<https://github.com/esimov/gogu/blob/master/generic.go#L8>)

CompFn is a generic function type for comparing two values.

```go
type CompFn[T any] func(a, b T) bool
```

### func [NewMemoizer](<https://github.com/esimov/gogu/blob/master/memoize.go#L19>)

```go
func NewMemoizer[T ~string, V any](expiration, cleanup time.Duration) *Memoizer[T, V]
```

NewMemoizer instantiates a new Memoizer.

### func \(Memoizer\[T, V\]\) [Memoize](<https://github.com/esimov/gogu/blob/master/memoize.go#L32>)

```go
func (m Memoizer[T, V]) Memoize(key T, fn func() (*cache.Item[V], error)) (*cache.Item[V], error)
```

Memoize returns the item under a specific key instantly in case the key exists, otherwise returns the results of the given function, making sure that only one execution is in\-flight for a given key at a time.

This method is useful for caching the result of a time\-consuming operation when is more important to return a slightly outdated result, than to wait for an operation to complete before serving it.

<details><summary>Example</summary>
<p>

```go
{
	m := NewMemoizer[string, any](time.Second, time.Minute)

	sampleItem := map[string]any{
		"foo": "one",
		"bar": "two",
		"baz": "three",
	}

	expensiveOp := func() (*cache.Item[any], error) {
		// Here we are simulating an expensive operation.
		time.Sleep(500 * time.Millisecond)

		foo := FindByKey(sampleItem, func(key string) bool {
			return key == "foo"
		})
		m.Cache.MapToCache(foo, cache.DefaultExpiration)

		item, err := m.Cache.Get("foo")
		if err != nil {
			return nil, err
		}
		return item, nil
	}

	fmt.Println(m.Cache.List())
	// Caching the result of some expensive fictive operation result.
	data, _ := m.Memoize("key1", expensiveOp)
	fmt.Println(len(m.Cache.List()))

	item, _ := m.Cache.Get("key1")
	fmt.Println(item.Val())

	// Serving the expensive operation result from the cache. This should return instantly.
	// If it would invoked the expensiveOp function this would be introduced a 500 millisecond latency.
	data, _ = m.Memoize("key1", expensiveOp)
	fmt.Println(data.Val())

}
```

#### Output

```
map[]
2
one
one
```

</p>
</details>

## type [Number](<https://github.com/esimov/gogu/blob/master/map.go#L12-L14>)

Number is a custom type set of constraints extending the Float and Integer type set from the experimental constraints package.

```go
type Number interface {
    // contains filtered or unexported methods
}
```

## type [RType](<https://github.com/esimov/gogu/blob/master/func.go#L68-L70>)

RType is a generic struct type used as method receiver on retry operations.

```go
type RType[T any] struct {
    Input T
}
```

### func \(RType\[T\]\) [Retry](<https://github.com/esimov/gogu/blob/master/func.go#L74>)

```go
func (v RType[T]) Retry(n int, fn func(T) error) (int, error)
```

Retry tries to invoke the callback function \`n\` times. It runs until the number of attempts is reached or the returned value of the callback function is nil.

<details><summary>Example</summary>
<p>

```go
{
	n := 2
	idx := 0
	ForEach([]string{"one", "two", "three"}, func(val string) {
		rt := RType[string]{Input: val}
		attempts, e := rt.Retry(n, func(elem string) (err error) {
			if len(elem)%3 != 0 {
				err = fmt.Errorf("retry failed: number of %d attempts exceeded", n)
			}
			return err
		})
		switch idx {
		case 0:
			fmt.Println(attempts)
		case 1:
			fmt.Println(attempts)
		case 2:
			fmt.Println(attempts)
			fmt.Println(e)
		}
		idx++
	})

}
```

#### Output

```
0
0
2
retry failed: number of 2 attempts exceeded
```

</p>
</details>

### func \(RType\[T\]\) [RetryWithDelay](<https://github.com/esimov/gogu/blob/master/func.go#L96>)

```go
func (v RType[T]) RetryWithDelay(n int, delay time.Duration, fn func(time.Duration, T) error) (time.Duration, int, error)
```

RetryWithDelay tries to invoke the callback function \`n\` times, but with a delay between each call. It runs until the number of attempts is reached or the error return value of the callback function is nil.

<details><summary>Example</summary>
<p>

```go
{
	n := 5
	// In this example we are simulating an external service. In case the response time
	// exceeds a certain time limit we stop retrying and we are returning an error.
	services := []struct {
		service string
		time    time.Duration
	}{
		{service: "AWS1"},
		{service: "AWS2"},
	}

	type Service[T ~string] struct {
		Service T
		Time    time.Duration
	}

	for _, srv := range services {
		r := random(1, 10)
		// Here we are simulating the response time of the external service
		// by generating some random duration between 1ms and 10ms.
		// All the test should pass because all of the responses are inside the predefined limit (10ms).
		service := Service[string]{
			Service: srv.service,
			Time:    time.Duration(r) * time.Millisecond,
		}
		rtyp := RType[Service[string]]{
			Input: service,
		}

		d, att, e := rtyp.RetryWithDelay(n, 20*time.Millisecond, func(d time.Duration, srv Service[string]) (err error) {
			if srv.Time.Milliseconds() > 10 {
				err = fmt.Errorf("retry failed: service time exceeded")
			}
			return err
		})
		fmt.Println(e)
		fmt.Println(att)
		fmt.Println(d.Milliseconds())
	}

}
```

#### Output

```
<nil>
0
0
```

</p>
</details>

## 🤝 Contributing
- Request new features or fix [open issues](https://github.com/esimov/gogu/issues)
- Fork the project and make [pull requests](https://github.com/esimov/gogu/pulls)

## Author
* Endre Simo ([@simo_endre](https://twitter.com/simo_endre))

## License
Copyright © 2022 Endre Simo

This software is distributed under the MIT license. See the [LICENSE](https://github.com/esimov/gogu/blob/master/LICENSE) file for the full license text.
