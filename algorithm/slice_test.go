package algorithm

import (
	"testing"
)

func TestUniqueAppend(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	itemA := 6
	itemB := 3
	if l := len(UniqueAppend(slice, itemA)); l != 6 {
		t.Errorf("len(UniqueAppend(%v, %v)) expect 6, got %v", slice, itemA, l)
	}
	if l := len(UniqueAppend(slice, itemB)); l != 5 {
		t.Errorf("len(UniqueAppend(%v, %v)) expect 5, got %v", slice, itemB, l)
	}
}
