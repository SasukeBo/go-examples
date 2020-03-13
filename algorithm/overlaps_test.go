package algorithm

import (
	"testing"
)

func TestOverlaps(t *testing.T) {
	l1 := []int{1, 2, 5, 6}
	l2 := []int{5, 7, 9}

	if !Overlaps(l1, l2) {
		t.Fatal("expect overlaps true, get false")
	}
}
