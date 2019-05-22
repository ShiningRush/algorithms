package main

import (
	"fmt"
)

type LFUCache struct {
	data       map[int]int
	count      *RecentCount
	curPayload int
	capacity   int
}

type RecentCount struct {
	key   int
	count int
	next  *RecentCount
}

func NewRecentCount(key int) *RecentCount {
	return &RecentCount{
		key:   key,
		count: 1,
	}
}

func Constructor(capacity int) LFUCache {
	cache := LFUCache{}
	cache.data = make(map[int]int)
	cache.capacity = capacity
	return cache
}

func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}

	if v, ok := this.data[key]; ok {
		this.addToIndex(key)
		return v
	}

	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	this.addToIndex(key)
	this.data[key] = value
}

func (this *LFUCache) addToIndex(key int) {
	if this.count == nil {
		new := NewRecentCount(key)
		this.count = new
		this.curPayload++
		return
	}

	if _, ok := this.data[key]; ok {
		this.accessExistNode(key)
	} else {
		if this.curPayload < this.capacity {
			this.addNewNode(key)
		} else {
			this.replaceNode(key)
		}
	}
}

func (this *LFUCache) accessExistNode(key int) {
	cur := this.count
	var last *RecentCount
loop1:
	for cur != nil {
		if cur.key == key {
			cur.count++
			next := cur.next
			if next == nil || next.count > cur.count {
				break loop1
			}

		loop2:
			for {
				if next.next == nil || (next.next.count > cur.count) {
					break loop2
				}
				next = next.next
			}

			if this.count == cur {
				this.count = cur.next
			}
			if last != nil {
				last.next = cur.next
			}
			cur.next = next.next
			next.next = cur
		}

		last = cur
		cur = cur.next
	}
}

func (this *LFUCache) addNewNode(key int) {
	new := NewRecentCount(key)

	if this.count.count > 1 {
		new.next = this.count
		this.count = new
	} else if this.count.count == 1 {
		cur := this.count
	loop:
		for cur != nil {
			if cur.next == nil {
				cur.next = new
				break loop
			}

			if cur.next.count > 1 {
				new.next = cur.next
				cur.next = new
				break loop
			}

			cur = cur.next
		}
	}

	this.curPayload++
}

func (this *LFUCache) replaceNode(key int) {
	delete(this.data, this.count.key)

	new := NewRecentCount(key)
	cur := this.count
loop:
	for {
		if cur.next == nil || cur.next.count > 1 {
			break loop
		}
		cur = cur.next
	}

	new.next = cur.next
	if cur == this.count {
		this.count = new

	} else {
		this.count = this.count.next
		cur.next = new
	}
}

func main() {
	c := Constructor(5)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)
	fmt.Println(c.Get(4))
	fmt.Println(c.Get(3))
	c.Put(5, 5)
	c.Put(6, 6)
	c.Put(7, 7)
	c.Put(8, 8)
	c.Put(9, 9)
	c.Put(10, 10)
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(1))
}
