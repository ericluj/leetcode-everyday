package leetcode

// LRU ...
func LRU(operators [][]int, k int) []int {
	res := []int{}
	l := InitLRUCache(k)
	for _, v := range operators {
		if v[0] == 1 {
			l.Put(v[1], v[2])
		} else if v[0] == 2 {
			x := l.Get(v[1])
			res = append(res, x)
		}
	}
	return res
}

// DoubleLinkedNode ...
type DoubleLinkedNode struct {
	Key  int
	Val  int
	Pre  *DoubleLinkedNode
	Next *DoubleLinkedNode
}

// initDoubleLinkedNode ...
func initDoubleLinkedNode(k, v int) *DoubleLinkedNode {
	return &DoubleLinkedNode{
		Key: k,
		Val: v,
	}
}

// LRUCache ...
type LRUCache struct {
	Size int
	Cap  int
	M    map[int]*DoubleLinkedNode
	Head *DoubleLinkedNode
	Tail *DoubleLinkedNode
}

// InitLRUCache ...
func InitLRUCache(cap int) LRUCache {
	l := LRUCache{
		Size: 0,
		Cap:  cap,
		M:    map[int]*DoubleLinkedNode{},
		Head: initDoubleLinkedNode(0, 0), //初始化两个无用节点作为头尾指针
		Tail: initDoubleLinkedNode(0, 0),
	}
	l.Head.Next = l.Tail
	l.Tail.Pre = l.Head
	return l
}

// AddToHead ...
func (l *LRUCache) AddToHead(node *DoubleLinkedNode) {
	node.Next = l.Head.Next
	node.Pre = l.Head
	l.Head.Next.Pre = node
	l.Head.Next = node
}

// RemoveNode ...
func (l *LRUCache) RemoveNode(node *DoubleLinkedNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

// MoveToHead ...
func (l *LRUCache) MoveToHead(node *DoubleLinkedNode) {
	l.RemoveNode(node)
	l.AddToHead(node)
}

// RemoveTail ...
func (l *LRUCache) RemoveTail() *DoubleLinkedNode {
	node := l.Tail.Pre
	l.RemoveNode(node)
	return node
}

// Get ...
func (l *LRUCache) Get(k int) int {
	if _, ok := l.M[k]; !ok {
		return -1
	}
	node := l.M[k]
	l.MoveToHead(node)
	return node.Val
}

// Put ...
func (l *LRUCache) Put(k, v int) {
	if _, ok := l.M[k]; !ok {
		node := initDoubleLinkedNode(k, v)
		l.M[k] = node
		l.AddToHead(node)
		l.Size++
		if l.Size > l.Cap {
			removed := l.RemoveTail()
			delete(l.M, removed.Key)
			l.Size--
		}
	} else {
		node := l.M[k]
		node.Val = v
		l.MoveToHead(node)
	}
}
