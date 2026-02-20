package student

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}
	i := 0
	for node := l; node != nil; node = node.Next {
		if i == pos {
			return node
		}
		i++
	}
	return nil
}
