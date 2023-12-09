package minstack

type MinStack struct {
	node *Node
}

type Node struct {
	Val  int
	Min  int
	Next *Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.node == nil {
		this.node = &Node{
			Val: val,
			Min: val,
		}
		return
	}

	this.node = &Node{
		Val:  val,
		Min:  min(this.node.Min, val),
		Next: this.node,
	}
	return
}

func (this *MinStack) Pop() {
	this.node = this.node.Next
}

func (this *MinStack) Top() int {
	return this.node.Val
}

func (this *MinStack) GetMin() int {
	return this.node.Min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
