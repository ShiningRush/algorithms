package philosopher

import (
	"log"
	"time"
	"math/rand"
)

type Philosopher struct{
	name string
	left *HierarchicalMutex
	right *HierarchicalMutex
	eatingCount int
}

func NewPhilosopher(name string,first *HierarchicalMutex,second *HierarchicalMutex) *Philosopher{
	n := &Philosopher{}
	n.name = name
	
	if first.level > second.level{
		n.left = first
		n.right = second
	} else {
		n.left = second
		n.right = first
	}


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