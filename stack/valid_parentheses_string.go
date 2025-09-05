package stack

// 若輸入的字串為合法字串，回傳原始字串
// 若非合法字串，將他擴展成一個合法字串
func convertToValidParenthesesString(s string) string {
	validStr := ""
	byteStack := new(Stack)
	for _, ch := range s {
		if _, ok := bracketsOpenCloseMap[byte(ch)]; ok {
			// open bracket
			byteStack.Push(byte(ch))
			validStr = validStr + string(ch)
		} else {
			// close bracket
			openBracket, popOk := byteStack.Pop()
			if popOk && bracketsOpenCloseMap[openBracket] == byte(ch) {
				validStr = validStr + string(ch)
			} else {
				byteStack.Push(openBracket)
				op := bracketsCloseOpenMap[byte(ch)]
				validStr = validStr + string(op)
				validStr = validStr + string(ch)
			}
		}
	}

	for !byteStack.IsEmpty() {
		openBracket, _ := byteStack.Pop()
		closeBracket := bracketsOpenCloseMap[openBracket]
		validStr = validStr + string(closeBracket)
	}
	return validStr
}

var bracketsCloseOpenMap = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}
