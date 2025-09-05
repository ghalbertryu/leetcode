package stack

import (
	"log"
	"testing"
)

func TestLinkedStack(t *testing.T) {
	s := new(LinkedStack)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	log.Println(s.IsEmpty())
	log.Println(s.Pop())
	log.Println(s.Pop())
	log.Println(s.Pop())
	log.Println(s.Pop())
	log.Println(s.IsEmpty())
}
