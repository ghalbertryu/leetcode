package leetcode

import (
	"testing"
)

func TestBuildLinkedList(t *testing.T) {
	nodeHead := buildLinkedList([]int{1, 2, 3})
	visitNodes(nodeHead)
}

var nodeHead1 = buildLinkedList([]int{1, 2, 4})
var nodeHead2 = buildLinkedList([]int{1, 3, 4, 8, 10})

func TestMergeTwoLists(t *testing.T) {
	mergedList := mergeTwoLists(nodeHead1, nodeHead2)
	visitNodes(mergedList)
}
