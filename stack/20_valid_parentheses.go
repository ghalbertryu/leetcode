package stack

func isValid(s string) bool {
	byteStack := new(Stack)
	for _, ch := range s {
		if _, ok := bracketsOpenCloseMap[byte(ch)]; ok {
			byteStack.Push(byte(ch))
		} else {
			pop, popOk := byteStack.Pop()
			if !popOk || bracketsOpenCloseMap[pop] != byte(ch) {
				return false
			}
		}
	}

	if !byteStack.IsEmpty() {
		return false
	}
	return true
}

var bracketsOpenCloseMap = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}
