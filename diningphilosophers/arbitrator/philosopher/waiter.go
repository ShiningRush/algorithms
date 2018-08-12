package philosopher

import (
	"sync"
)

var waiter Waiter

func init(){
	waiter = Waiter{}
}

type Waiter struct{
	sync.Mutex
}

func(w *Waiter) GetPermission(){
	waiter.Lock()
}

func(w *Waiter) Done(){
	waiter.Unlock()
}