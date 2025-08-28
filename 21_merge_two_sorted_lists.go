package leetcode

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildLinkedList(sortedNums []int) *ListNode {
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

func visitNodes(nodeHead *ListNode) {
	log.Println(nodeHead.Val)
	for nodeHead.Next != nil {
		nodeHead = nodeHead.Next
		log.Println(nodeHead.Val)
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	tmpNodeList1 := list1
	tmpNodeList2 := list2
	var retNode *ListNode
	var tmpNode *ListNode
	for tmpNodeList1 != nil && tmpNodeList2 != nil {
		if tmpNodeList1.Val <= tmpNodeList2.Val {
			tmpNode.Next = tmpNodeList1
			tmpNodeList1 = tmpNodeList1.Next
		} else {
			tmpNode.Next = tmpNodeList2
			tmpNodeList2 = tmpNodeList2.Next
		}
		if retNode == nil {
			retNode = tmpNode
		}
	}

	return retNode
}
