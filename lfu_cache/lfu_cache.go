package lfu_cache

import "container/list"

type Tuple struct {
	value  int
	parent *list.Element
	node   *list.Element
}

type Entry struct {
	value int
	list  *list.List
}

type LFUCache struct {
	freq  *list.List
	cache map[int]*Tuple
	cap   int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{freq: list.New(),
		cache: map[int]*Tuple{},
		cap:   capacity,
	}
}

func (this *LFUCache) Update(key int) {
	tuple := this.cache[key]
	freq := tuple.parent
	entry := freq.Value.(Entry)
	nextFreq := freq.Next()

	if nextFreq == nil || entry.value+1 != nextFreq.Value.(Entry).value {
		tmp := Entry{entry.value + 1, list.New()}
		nextFreq = this.freq.InsertAfter(tmp, freq)
	}

	entry.list.Remove(tuple.node)
	if entry.list.Len() == 0 {
		this.freq.Remove(freq)
	}

	nextEntry := nextFreq.Value.(Entry)
	tuple.parent = nextFreq
	tuple.node = nextEntry.list.PushFront(key)
}

func (this *LFUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.Update(key)
		return elem.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.Update(key)
		elem.value = value
	} else if this.cap > 0 {
		if len(this.cache) >= this.cap {
			entry := this.freq.Front().Value.(Entry)
			delKey := entry.list.Back().Value.(int)
			delete(this.cache, delKey)
			entry.list.Remove(entry.list.Back())
			if entry.list.Len() == 0 {
				this.freq.Remove(this.freq.Front())
			}
		}

		if this.freq.Len() == 0 || this.freq.Front().Value.(Entry).value != 1 {
			tmp := Entry{1, list.New()}
			this.freq.PushFront(tmp)
		}

		entry := this.freq.Front().Value.(Entry)
		tmp := new(Tuple)
		tmp.value = value
		tmp.node = entry.list.PushFront(key)
		tmp.parent = this.freq.Front()
		this.cache[key] = tmp
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
