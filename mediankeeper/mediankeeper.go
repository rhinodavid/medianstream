package mediankeeper

import (
	"container/heap"

	"github.com/rhinodavid/medianstream/intheap"
)

// MedianKeeper is a structure designed to find the median of a series of
// integers in an efficient manner.
type MedianKeeper struct {
	l, h *intheap.IntHeap
}

// New creates and intializes a new MedianKeeper
func New() *MedianKeeper {
	var k MedianKeeper
	k.l = &intheap.IntHeap{}
	k.h = &intheap.IntHeap{}
	heap.Init(k.l)
	heap.Init(k.h)
	return &k
}

// Push adds a new integer to the structure. Runs in logarithmic time
// with respect to the total number of integers added
func (k *MedianKeeper) Push(x int) {
	if k.l.Len() == 0 && k.h.Len() == 0 {
		// first item
		heap.Push(k.l, -1*x)
	} else if k.l.Len() == 0 {
		// second item
		f := heap.Pop(k.l).(int) * -1
		if f < x {
			heap.Push(k.h, x)
			heap.Push(k.l, f*-1)
		} else {
			// swap
			heap.Push(k.l, x*-1)
			heap.Push(k.h, f)
		}
	} else {
		if k.l.Len() > k.h.Len() {
			lV := k.l.Peek().(int) * -1
			if x >= lV {
				heap.Push(k.h, x)
			} else {
				lI := heap.Pop(k.l).(int) * -1
				heap.Push(k.h, lI)
				heap.Push(k.l, x*-1)
			}
		} else {
			hV := k.h.Peek().(int)
			if x <= hV {
				heap.Push(k.l, x*-1)
			} else if x > hV {
				hI := heap.Pop(k.h).(int)
				heap.Push(k.l, hI*-1)
				heap.Push(k.h, x)
			}
		}
	}
}

// Median returns the median of the integers inserted into the structure.
// If there is an even number of integers, it returns the smaller of the
// two medians. Runs in constant time with respect to the total number of
// integers inserted.
func (k *MedianKeeper) Median() int {
	return k.l.Peek().(int) * -1
}
