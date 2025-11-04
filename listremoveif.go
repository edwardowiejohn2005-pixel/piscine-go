package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	// Remove nodes from the beginning if they match
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// If the list became empty after removals
	if l.Head == nil {
		l.Tail = nil
		return
	}

	prev := l.Head
	current := l.Head.Next

	for current != nil {
		if current.Data == data_ref {
			prev.Next = current.Next
			if current == l.Tail {
				l.Tail = prev
			}
		} else {
			prev = current
		}
		current = current.Next
	}
}
