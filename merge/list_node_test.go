package merge

import "testing"

func TestBuildLinkedList(t *testing.T) {
	nodeHead := NewListNode([]int{1, 2, 3})
	VisitList(nodeHead)
}
