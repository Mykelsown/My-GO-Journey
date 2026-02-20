package student

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	var result *NodeI

	if n1.Data <= n2.Data {
		result = n1
		result.Next = SortedListMerge(n1.Next, n2)
	} else {
		result = n2
		result.Next = SortedListMerge(n1, n2.Next)
	}

	return result
}
