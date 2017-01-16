package mediankeeper

import (
	"container/heap"

	"github.com/rhinodavid/medianstream/intheap"
)

type MedianKeeper struct {
	l, h *intheap.IntHeap
}

func New() *MedianKeeper {
	var k MedianKeeper
	k.l = &intheap.IntHeap{}
	k.h = &intheap.IntHeap{}
	heap.Init(k.l)
	heap.Init(k.h)
	return &k
}

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

func (k *MedianKeeper) Median() int {
	return k.l.Peek().(int) * -1
}
