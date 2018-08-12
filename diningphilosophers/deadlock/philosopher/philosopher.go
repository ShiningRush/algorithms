package philosopher

import (
	"sync"
	"log"
	"time"
	"math/rand"
)

type Philosopher struct{
	name string
	left *sync.Mutex
	right *sync.Mutex
	eatingCount int
}

func NewPhilosopher(name string,left *sync.Mutex,right *sync.Mutex) *Philosopher{
	n := &Philosopher{}
	n.name = name
	n.left = left
	n.right = right

	return n
}

func(p *Philosopher) Start(){
	for true{
		p.takeRest()
		p.haveEating()
	}
}

func(p *Philosopher) takeRest(){
	log.Println(p.name + " have a rest")
	p.watingRandonTime()
}

func(p *Philosopher) haveEating(){
	log.Println(p.name + " take to left")
	p.left.Lock()
	p.watingRandonTime() // get some delay, you will see deadlock soon
	log.Println(p.name + " take to right")
	p.right.Lock()
	log.Println(p.name + " ready to eating")
	p.eatingCount++
	log.Println(p.name + " put right")
	p.right.Unlock()
	log.Println(p.name + " put left")
	p.left.Unlock()
}

func (p *Philosopher) watingRandonTime(){
	ri := rand.Int() % 200
	time.Sleep(time.Duration(ri) * time.Millisecond)
}