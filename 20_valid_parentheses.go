package leetcode

func isValid(s string) bool {
	filo := new(FiLo)
	for _, ch := range s {
		if _, ok := signPairMap[byte(ch)]; ok {
			filo.Push(byte(ch))
		} else {
			pop := filo.Pop()
			if signPairMap[pop] != byte(ch) {
				return false
			}
		}
	}

	if filo.Len() != 0 {
		return false
	}
	return true
}

var signPairMap = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

type FiLo []byte

func (o *FiLo) Push(b byte) {
	*o = append(*o, b)
}

func (o *FiLo) Pop() byte {
	l := len(*o)
	if l == 0 {
		return 0
	}

	ret := (*o)[l-1]
	(*o) = (*o)[:l-1]
	return ret
}

func (o *FiLo) Len() int {
	return len(*o)
}
