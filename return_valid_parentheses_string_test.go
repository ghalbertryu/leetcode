package leetcode

import (
	"log"
	"testing"
)

func TestReturnValidParenthesesString(t *testing.T) {
	log.Println(returnValidParenthesesString("()"))
	log.Println(returnValidParenthesesString("()[]{}"))
	log.Println(returnValidParenthesesString("([])"))

	log.Println(returnValidParenthesesString("[()")) // 沒關
	log.Println(returnValidParenthesesString("(]"))  // 沒開
}
