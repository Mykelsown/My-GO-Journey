package student

func ListPushFront(l *List, data interface{}) {
	if l == nil {
		return
	}

	newNode := &NodeL{
		Data: data,
		Next: l.Head,
	}

	l.Head = newNode

	if l.Tail == nil {
		l.Tail = newNode
	}
}
