package lrucache

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Post *Node
}

type LRUCache struct {
	Count map[int]*Node
	Head  *Node
	Tail  *Node

	capacity int
}

func Constructor(capacity int) LRUCache {
	headNode := &Node{}
	tailNode := &Node{}

	headNode.Post = tailNode
	tailNode.Pre = headNode

	return LRUCache{
		Count:    make(map[int]*Node, capacity),
		Head:     headNode,
		Tail:     tailNode,
		capacity: capacity,
	}
}

func (this *LRUCache) moveToTail(node *Node) {
	nodePre := node.Pre
	nodePost := node.Post
	if nodePre != nil {
		nodePre.Post = nodePost
	}
	if nodePost != nil {
		nodePost.Pre = nodePre
	}

	this.Tail.Pre.Post = node
	node.Pre = this.Tail.Pre
	this.Tail.Pre = node
	node.Post = this.Tail
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Count[key]
	if !ok {
		return -1
	}

	this.moveToTail(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Count[key]
	if !ok {
		if len(this.Count) == this.capacity {
			// remove head node.
			headNode := this.Head.Post
			this.Head.Post = headNode.Post
			headNode.Post.Pre = this.Head
			delete(this.Count, headNode.Key)
		}

		newNode := &Node{Key: key, Val: value}

		this.Count[key] = newNode
		this.moveToTail(newNode)
		return
	}
	node.Val = value
	this.moveToTail(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
