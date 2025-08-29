package merge

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) < 1 {
		return nil
	}

	for len(lists) > 1 {
		tmpLists := make([]*ListNode, 0)
		for i := 0; i < len(lists); i = i + 2 {
			list1 := lists[i]
			if i+1 == len(lists) {
				tmpLists = append(tmpLists, list1)
				break
			}

			list2 := lists[i+1]
			tmpLists = append(tmpLists, mergeTwoLists(list1, list2))
		}
		lists = tmpLists
	}

	return lists[0]
}
