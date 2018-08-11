package philosopher

import (
	"sync"
)

type Status uint
const (
	Dirty  Status = iota
	Clean
)

type Resource struct{
	Status
	sync.Mutex
}

func(r *Resource) IsClean() bool{
	return r.Status == Clean
}