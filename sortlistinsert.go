package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	node := &NodeI{Data: data_ref}

	// Insert at the beginning if the list is empty or data_ref should be the new head
	if l == nil || data_ref < l.Data {
		node.Next = l
		return node
	}

	// Find insertion point (after nodes with smaller Data)
	prev := l
	for prev.Next != nil && prev.Next.Data < data_ref {
		prev = prev.Next
	}

	// Insert node
	node.Next = prev.Next
	prev.Next = node

	return l
}
