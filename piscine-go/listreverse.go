package student

func ListReverse(l *List) {
	if l == nil || l.Head == nil || l.Head.Next == nil {
		return
	}
	var prev *NodeL
	curr := l.Head
	l.Tail = l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}
