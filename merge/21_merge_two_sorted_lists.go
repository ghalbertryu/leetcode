package merge

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// init retNode and tmpNode
	var retNode *ListNode
	var tmpNode *ListNode
	if list1.Val <= list2.Val {
		retNode = list1
		tmpNode = list1
		list1 = list1.Next
	} else {
		retNode = list2
		tmpNode = list2
		list2 = list2.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tmpNode.Next = list1
			list1 = list1.Next
		} else {
			tmpNode.Next = list2
			list2 = list2.Next
		}
		tmpNode = tmpNode.Next
	}

	// 連接剩餘 list
	if list1 != nil {
		tmpNode.Next = list1
	}
	if list2 != nil {
		tmpNode.Next = list2
	}

	return retNode
}
