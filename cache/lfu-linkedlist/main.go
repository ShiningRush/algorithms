package main

import (
	"fmt"
)

type LFUCache struct {
	data   map[int]int
	count  []RecentCount
	access int
}

type RecentCount struct {
	key    int
	count  int
	access int
}

func (rc RecentCount) Access(acc int) {
	rc.access = acc
}

func Constructor(capacity int) LFUCache {
	cache := LFUCache{}
	cache.data = make(map[int]int)
	cache.count = make([]RecentCount, capacity, capacity)
	return cache
}

func (this *LFUCache) Get(key int) int {
	if len(this.count) == 0 {
		return -1
	}

	if v, ok := this.data[key]; ok {
		this.addToIndex(key)
		return v
	}

	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if len(this.count) == 0 {
		return
	}

	this.data[key] = value
	this.addToIndex(key)
}

func (this *LFUCache) addToIndex(key int) {
	min := this.count[0].count
	minIdx := 0
	minTime := this.count[0].access
	findFlg := false
loop:
	for idx, v := range this.count {
		if v.count == 0 {
			this.count[idx].key = key
			this.count[idx].count++
			this.count[idx].access = this.access
			findFlg = true
			break loop
		}

		if v.key == key {
			this.count[idx].count++
			this.count[idx].access = this.access
			findFlg = true
			break loop
		}

		if v.count < min {
			min = v.count
			minIdx = idx
			minTime = v.access
		}

		if v.count == min && v.access < minTime {
			min = v.count
			minIdx = idx
			minTime = v.access
		}
	}

	if !findFlg {
		delete(this.data, this.count[minIdx].key)

		this.count[minIdx].key = key
		this.count[minIdx].count = 1
		this.count[minIdx].access = this.access
	}

	this.access++
}

func main() {
	c := Constructor(2)
	c.Put(2, 1)
	c.Put(3, 2)
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(2))
	c.Put(4, 3)
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
}
