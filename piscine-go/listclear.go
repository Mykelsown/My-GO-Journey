package student

func ListClear(l *List) {
	if l == nil {
		return
	}
	*l = List{}
}
