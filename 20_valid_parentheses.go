package leetcode

func isValid(s string) bool {
	byteStack := new(Stack)
	for _, ch := range s {
		if _, ok := bracketsPairMap[byte(ch)]; ok {
			byteStack.Push(byte(ch))
		} else {
			pop, popOk := byteStack.Pop()
			if popOk && bracketsPairMap[pop] != byte(ch) {
				return false
			}
		}
	}

	if byteStack.IsEmpty() {
		return false
	}
	return true
}

var bracketsPairMap = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

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
