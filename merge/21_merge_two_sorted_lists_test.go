package merge

import (
	"testing"
)

var nodeHead1 = NewListNode([]int{1, 2, 4})
var nodeHead2 = NewListNode([]int{1, 3, 4, 8, 10})

func TestMergeTwoLists(t *testing.T) {
	mergedList := mergeTwoLists(nodeHead1, nodeHead2)
	VisitList(mergedList)
}
