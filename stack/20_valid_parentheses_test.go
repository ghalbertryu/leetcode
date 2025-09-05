package stack

import (
	"log"
	"testing"
)

func TestIsValid(t *testing.T) {
	log.Println(isValid("()"))
	log.Println(isValid("()[]{}"))
	log.Println(isValid("(]"))
	log.Println(isValid("([])"))
}
