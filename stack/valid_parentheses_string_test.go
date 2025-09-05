package stack

import (
	"log"
	"testing"
)

func TestConvertToValidParenthesesString(t *testing.T) {
	log.Println(convertToValidParenthesesString("()"))
	log.Println(convertToValidParenthesesString("()[]{}"))
	log.Println(convertToValidParenthesesString("([])"))

	log.Println(convertToValidParenthesesString("[()")) // 沒關
	log.Println(convertToValidParenthesesString("(]"))  // 沒開
}
