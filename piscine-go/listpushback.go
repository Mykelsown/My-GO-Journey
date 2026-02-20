package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	if l == nil {
		return
	}
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	if l.Tail == nil {
		cur := l.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = n
		l.Tail = n
		return
	}

	l.Tail.Next = n
	l.Tail = n
}
