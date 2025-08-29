package merge

import (
	"container/heap"
)

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

func mergeKListsUsingHeap(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) < 1 {
		return nil
	}

	pq := &ListNodeHeap{}
	for _, node := range lists {
		if node != nil {
			heap.Push(pq, node)
		}
	}

	var head *ListNode
	var tmpNode *ListNode
	listNode := heap.Pop(pq).(*ListNode)
	head = listNode
	tmpNode = listNode
	if listNode.Next != nil {
		listNode = listNode.Next
		heap.Push(pq, listNode)
	}

	for pq.Len() > 0 {
		popNode := heap.Pop(pq).(*ListNode)
		tmpNode.Next = popNode
		if pq.Len() == 0 {
			return head
		}

		tmpNode = tmpNode.Next
		if popNode.Next != nil {
			popNode = popNode.Next
			heap.Push(pq, popNode)
		}
	}
	return head
}

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int {
	return len(h)
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h *ListNodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
