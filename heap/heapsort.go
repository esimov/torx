package heap

import (
	"sync"

	"github.com/esimov/gogu"
)

// Sort sorts the heap in ascending or descening order, depending on the heap type.
// If the heap is a max heap, the heap is sorted in ascending order,
// otherwise if the heap is a min heap, it is sorted in descending order.
func Sort[T comparable](mu *sync.RWMutex, data []T, comp gogu.CompFn[T]) []T {
	heap := FromSlice(mu, data, comp)

	for i := heap.Size() - 1; i > 0; i-- {
		swap(mu, data, 0, i)
		heap.moveDown(i, 0)
	}

	return heap.GetValues()
}
