package cache

type node struct {
	next *node
	prev *node
	val  int
}

type LRUCache struct {
	Capacity int
	Curr     int
	Map      map[int]*node
	Head     *node // head points to most recently used
	Tail     *node // tail points to least recently used
}

func Constructor(capacity int) *LRUCache {
	// Initialize nodes that point to the most and least recently used nodes
	head := &node{
		next: nil,
		prev: nil,
		val:  0,
	}
	tail := &node{
		next: nil,
		prev: nil,
		val:  0,
	}
	// Before that they'll point to each other because why not
	head.next = tail
	tail.prev = head

	// Initialize cache
	cache := &LRUCache{
		Capacity: capacity,
		Map:      map[int]*node{},
		Head:     head,
		Tail:     tail,
		Curr:     0,
	}

	return cache
}

func (this *LRUCache) insertNode(node *node) {
	headsNext := this.Head.next
	this.Head.next = node
	node.prev = this.Head
	node.next = headsNext
	headsNext.prev = node
}

func (this *LRUCache) removeNode(node *node) {
	next, prev := node.next, node.prev
	prev.next = next
	next.prev = prev
	node = nil
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.Map[key]; ok {
		this.removeNode(val)
		this.insertNode(val)
		return val.val
	}
	return -1
}

func (this *LRUCache) Put(key int, val int) {
	insertMe := &node{
		val:  val,
		next: nil,
		prev: nil,
	}
	this.insertNode(insertMe)
	this.Map[key] = insertMe
	this.Curr++

	if this.Curr > this.Capacity {
		this.removeNode(this.Tail.prev)
	}
}
