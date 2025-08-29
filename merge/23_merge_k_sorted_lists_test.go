package merge

import (
	"testing"
)

var list1 = NewListNode([]int{1, 2, 4})
var list2 = NewListNode([]int{1, 3, 4, 8, 10})
var list3 = NewListNode([]int{1, 3, 4, 5})

func TestMergeKLists(t *testing.T) {
	mergedList := mergeKLists([]*ListNode{
		list1,
		list2,
		list3,
	})
	VisitList(mergedList)
}

func TestMergeKListsUsingHeap(t *testing.T) {
	mergedList := mergeKListsUsingHeap([]*ListNode{
		list1,
		list2,
		list3,
	})
	VisitList(mergedList)
}
