package stack

type LinkedStack struct {
	head *LinkedNode
}

type LinkedNode struct {
	Value int32
	Next  *LinkedNode
}

func (o *LinkedStack) Push(b byte) {
	newHeadNode := &LinkedNode{
		Value: int32(b),
		Next:  o.head,
	}
	o.head = newHeadNode
}

func (o *LinkedStack) Pop() (byte, bool) {
	if o.head == nil {
		return 0, false
	}

	val := o.head.Value
	o.head = o.head.Next
	return byte(val), true
}

func (o *LinkedStack) IsEmpty() bool {
	return o.head == nil
}
