package main

import (
	"fmt"
)

type LFUCache struct {
	data       map[int]*DataNode
	freq       map[int]*RecentCount
	curPayload int
	capacity   int
	minFreq    int
}

type RecentCount struct {
	key  int
	next *RecentCount
}

type DataNode struct {
	count int
	data  int
}

func NewRecentCount(key int) *RecentCount {
	return &RecentCount{
		key: key,
	}
}

func (rc *RecentCount) Add(newNode *RecentCount) {
	next := rc.next
	cur := rc
	for next != nil {
		cur = next
		next = next.next
	}

	cur.next = newNode
}

func (rc *RecentCount) Remove(key int) {
	cur := rc.next
	last := rc
	for cur != nil {
		if cur.key == key {
			last.next = cur.next
		}

		last = cur
		cur = cur.next
	}
}

func Constructor(capacity int) LFUCache {
	cache := LFUCache{}
	cache.data = make(map[int]*DataNode)
	cache.freq = make(map[int]*RecentCount)
	cache.capacity = capacity
	return cache
}

func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}

	if v, ok := this.data[key]; ok {
		v.count++
		this.refreshFreq(key)
		return v.data
	}

	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	if v, ok := this.data[key]; ok {
		v.count++
	} else {
		this.data[key] = &DataNode{
			count: 1,
			data:  value,
		}
	}
	this.refreshFreq(key)
}

func (this *LFUCache) refreshFreq(key int) {
	d := this.data[key]
	newNode := NewRecentCount(key)
	if this.freq[d.count] == nil {
		this.freq[d.count] = newNode
		this.curPayload++
		return
	}

	cur := this.freq[d.count]
	var last *RecentCount
	for cur != nil {
		if cur.key == key {
			if cur == this.freq[d.count] {
				this.freq[d.count] = cur.next
			} else {
				last.next = cur.next
			}

			return
		}

		last = cur
		cur = cur.next
	}

	if this.curPayload < this.capacity {
		last.next = newNode
	} else {
		newNode.next = this.freq[d.count].next
		this.freq[d.count] = newNode
	}
	this.curPayload++
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
