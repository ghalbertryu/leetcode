package stack

type Stack []byte

func (o *Stack) Push(b byte) {
	*o = append(*o, b)
}

func (o *Stack) Pop() (byte, bool) {
	l := len(*o)
	if l == 0 {
		return 0, false
	}

	ret := (*o)[l-1]
	(*o) = (*o)[:l-1]
	return ret, true
}

func (o *Stack) IsEmpty() bool {
	return len(*o) == 0
}
