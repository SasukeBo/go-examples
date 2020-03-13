package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a := l1.Val + l2.Val
	node := &ListNode{Val: a % 10}
	a = a / 10
	out := node
	l1 = l1.Next
	l2 = l2.Next

	for {
		if l1 == nil && l2 == nil {
			if a != 0 {
				node.Next = &ListNode{Val: a}
			}
			return out
		} else if l1 == nil {
			a = l2.Val + a
		} else if l2 == nil {
			a = l1.Val + a
		} else {
			a = l1.Val + l2.Val + a
		}
		node.Next = &ListNode{Val: a % 10}
		node = node.Next
		a = a / 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	return addNode(l1, l2, 0)
}

func addNode(l1 *ListNode, l2 *ListNode, pre int) *ListNode {
	if l1 == nil && l2 == nil {
		if pre != 0 {
			return &ListNode{Val: pre}
		} else {
			return nil
		}
	}
	if l1 == nil {
		l1 = &ListNode{Val: 0}
	}
	if l2 == nil {
		l2 = &ListNode{Val: 0}
	}
	totalSum := l1.Val + l2.Val + pre
	i, j := totalSum/10, totalSum%10
	l := new(ListNode)
	l.Val = j
	l.Next = addNode(l1.Next, l2.Next, i)
	return l
}
