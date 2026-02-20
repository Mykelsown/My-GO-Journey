package student

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	if l == nil || f == nil || cond == nil {
		return
	}
	for node := l.Head; node != nil; node = node.Next {
		if cond(node) {
			f(node)
		}
	}
}

func IsPositiveNode(node *NodeL) bool {
	if node == nil {
		return false
	}
	switch v := node.Data.(type) {
	case int:
		return v > 0
	case float32:
		return v > 0
	case float64:
		return v > 0
	case byte:
		return v > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	if node == nil {
		return false
	}
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
