package pqueue

import (
	"container/heap"
)

// PriorityQueue是Item指针的切片，实现一个最小堆

type Item struct {
	Value    interface{}   //可以存任意类型
	Priority int64
	Index    int
}

// this is a priority queue as implemented by a min heap
// ie. the 0th element is the *lowest* value
type PriorityQueue []*Item

//创建这个数组，capacity是容量，初始大小是0
func New(capacity int) PriorityQueue {
	return make(PriorityQueue, 0, capacity)
}

//返回数组的长度
func (pq PriorityQueue) Len() int {
	return len(pq)
}
//比较两个元素的优先级
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

//交换两个切片的元素
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

//将x加入堆中
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	c := cap(*pq)
	//如果加入一个元素后超过容量了，那么手动分配内存
	if n+1 > c {
		//手动分配内存，长度是n，容量翻倍
		npq := make(PriorityQueue, n, c*2)
		copy(npq, *pq)
		*pq = npq
	}
	//扩容pq的长度加1
	*pq = (*pq)[0 : n+1]
	//断言x的类型
	item := x.(*Item)
	item.Index = n
	(*pq)[n] = item //添加这个元素
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	c := cap(*pq)
	//缩小数组的capacity
	if n < (c/2) && c > 25 {
		npq := make(PriorityQueue, n, c/2)
		copy(npq, *pq)
		*pq = npq
	}
	item := (*pq)[n-1]
	item.Index = -1
	*pq = (*pq)[0 : n-1]
	return item
}

//返回堆顶元素，如果参数max小于堆顶元素优先级，返回的是负数
func (pq *PriorityQueue) PeekAndShift(max int64) (*Item, int64) {
	if pq.Len() == 0 {
		return nil, 0
	}

	item := (*pq)[0]
	if item.Priority > max {
		return nil, item.Priority - max
	}
	heap.Remove(pq, 0)

	return item, 0
}
