package topkfrequentelements

import "container/heap"

type Item struct {
	Num  int
	Freq int
}

type MaxHeap []Item

func (mh MaxHeap) Len() int           { return len(mh) }
func (mh MaxHeap) Less(i, j int) bool { return mh[i].Freq > mh[j].Freq }
func (mh MaxHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MaxHeap) Push(item any)     { *mh = append(*mh, item.(Item)) }
func (mh *MaxHeap) Pop() any {
	item := (*mh)[mh.Len()-1]
	*mh = (*mh)[:mh.Len()-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	// O(n)
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// O(n)
	mh := make(MaxHeap, 0, len(freqMap))
	for key, val := range freqMap {
		mh = append(mh, Item{Num: key, Freq: val})
	}

	ans := make([]int, 0, k)
	heap.Init(&mh) // O(n)
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(&mh).(Item).Num)
	}

	return ans
}
