package student

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	dummy := &NodeI{Next: l}
	prev := l
	current := l.Next

	for current != nil {
		if prev.Data <= current.Data {
			prev = current
			current = current.Next
		} else {
			prev.Next = current.Next

			var startPrev *NodeI = dummy
			start := dummy.Next

			for start.Data < current.Data {
				startPrev = start
				start = start.Next
			}

			startPrev.Next = current
			current.Next = start

			current = prev.Next
		}
	}

	return dummy.Next
}
