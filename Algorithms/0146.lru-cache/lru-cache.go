package problem0146

import "container/list"

type Pair struct {
	key   int
	value int
}
type LRUCache struct {
	list  *list.List
	cache map[int]*list.Element
	cap   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cap: capacity,
		list:  list.New(),
		cache: make(map[int]*list.Element)}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		if this.list.Len() >= this.cap {
			delete(this.cache, this.list.Back().Value.(Pair).key)
			this.list.Remove(this.list.Back())
		}
		this.list.PushFront(Pair{key, value})
		this.cache[key] = this.list.Front()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
