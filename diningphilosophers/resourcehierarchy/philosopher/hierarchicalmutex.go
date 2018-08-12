package philosopher

import (
	"sync"
)

type HierarchicalMutex struct{
	level int
	sync.Mutex
}

func NewHierarchicalMutex(level int) *HierarchicalMutex {
	n := &HierarchicalMutex{}
	n.level = level

	return n
}