package leetcode

// 若輸入的字串為合法字串，回傳原始字串
// 若非合法字串，將他擴展成一個合法字串
func returnValidParenthesesString(s string) string {
	validStr := ""
	byteStack := new(Stack)
	for _, ch := range s {
		if _, ok := bracketsPairMap[byte(ch)]; ok {
			// open bracket
			byteStack.Push(byte(ch))
			validStr = validStr + string(ch)
		} else {
			// close bracket
			openBracket, popOk := byteStack.Pop()
			if popOk && bracketsPairMap[openBracket] == byte(ch) {
				validStr = validStr + string(ch)
			} else {
				byteStack.Push(openBracket)
				op := bracketsPairMap2[byte(ch)]
				validStr = validStr + string(op)
				validStr = validStr + string(ch)
			}
		}
	}

	for !byteStack.IsEmpty() {
		openBracket, _ := byteStack.Pop()
		closeBracket := bracketsPairMap[openBracket]
		validStr = validStr + string(closeBracket)
	}
	return validStr
}

var bracketsPairMap2 = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}
