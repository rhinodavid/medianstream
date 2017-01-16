package mediankeeper

import (
	"testing"

	"github.com/rhinodavid/medianstream/mediankeeper"
)

func TestMedianstream(t *testing.T) {
	k := mediankeeper.New()
	// 1, 4, 6, -1, 45, 1

	nums := []int{1, 4, 6, -1, 45, 1, 100, 100, 84}
	medians := []int{1, 1, 4, 1, 4, 1, 4, 4, 6}

	for i, v := range nums {
		k.Push(v)
		if m := k.Median(); m != medians[i] {
			t.Errorf("Pushed %d, expected median %d, got %d\n", v, medians[i], m)
		}
	}
}
