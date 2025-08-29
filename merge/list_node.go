package merge

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(sortedNums []int) *ListNode {
	var tmpNodePtr *ListNode // default is nil
	for i := len(sortedNums) - 1; i >= 0; i-- {
		nodePtr := &ListNode{
			Val:  sortedNums[i],
			Next: tmpNodePtr,
		}
		tmpNodePtr = nodePtr
	}

	return tmpNodePtr
}

func VisitList(nodeHead *ListNode) {
	if nodeHead == nil {
		log.Println(nil)
		return
	}

	log.Println(nodeHead.Val)
	for nodeHead.Next != nil {
		nodeHead = nodeHead.Next
		log.Println(nodeHead.Val)
	}
}
