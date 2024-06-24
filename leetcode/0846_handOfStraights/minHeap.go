package handofstraights

type MinHeap []int

func (this MinHeap) Len() int {
	return len(this)
}

func (this MinHeap) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this MinHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *MinHeap) Push(item any) {
	*this = append(*this, item.(int))
}

func (this *MinHeap) Pop() any {
	old := *this
	item := old[len(old)-1]
	*this = old[:len(old)-1]
	return item
}
