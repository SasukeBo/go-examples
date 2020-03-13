package algorithm

import (
	"testing"
)

func TestAddTwoNumbers(b *testing.T) {
	l1 := &ListNode{
		Val: 0,
	}

	l2 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 3,
		},
	}

	addTwoNumbers(l1, l2)

	// node :=
	// result := make([]int, 0)
	// for {
	// if node != nil {
	// result = append(result, node.Val)
	// node = node.Next
	// } else {
	// break
	// }
	// }

	// if fmt.Sprintf("%v", result) != "[7 3]" {
	// t.Errorf("Expect [7 3], got %v\n", result)
	// }
}
