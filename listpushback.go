package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil { // if the list is empty
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode // attach new node after current tail
		l.Tail = newNode      // update tail to the new node
	}
}
