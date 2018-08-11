package philosopher

import (
	"log"
	"time"
	"sync"
	"math/rand"
)

type Philosopher struct{
	name string
	id int
	left *Resource
	right *Resource
	leftContending chan chan *Resource
	rightContending chan chan *Resource
	resChan chan *Resource
	eatingCount int
	isBusy sync.Mutex
}

func NewPhilosopher(
	name string, 
	id int, 
	left *Resource, right *Resource, 
	leftContending chan chan *Resource, rightContending chan chan *Resource) *Philosopher{
	n := &Philosopher{}
	n.name = name
	n.left = left
	n.right = right
	n.leftContending = leftContending
	n.rightContending = rightContending
	n.resChan = make(chan *Resource)

	return n
}

func(p *Philosopher) Start(){
	go func(){
		for{
			p.handRequest()
		}
	}()

	for{
		p.takeRest()
		p.haveEating()
		p.handRequest()
	}
}

func(p *Philosopher) takeRest(){
	log.Println(p.name + " have a rest")
	p.watingRandonTime()
}

func(p *Philosopher) haveEating(){
	p.isBusy.Lock()
	p.getLeft()
	p.watingRandonTime() // get some delay, you will see deadlock soon
	p.getRight()
	log.Println(p.name + " ready to eating")
	p.eatingCount++
	p.left.Status = Dirty
	p.right.Status = Dirty
	p.isBusy.Unlock()
}

func(p *Philosopher) getLeft(){
	if p.left == nil{
		log.Println(p.name + " has no left res, ready to get")
		p.leftContending <- p.resChan
		p.left = <- p.resChan
	}
	log.Println(p.name + " already get left res")
}

func(p *Philosopher) getRight(){
	if p.right == nil{
		log.Println(p.name + " has no right res, ready to get")
		p.rightContending <- p.resChan
		p.right = <- p.resChan
	}
	log.Println(p.name + " already get right res")
}

func(p *Philosopher) handRequest(){
	select{
	case leftr := <- p.leftContending: // put left
		if p.left != nil && !p.left.IsClean(){
			p.isBusy.Lock()
			log.Println(p.name + " get message, ready to pass left res")
			p.left.Status = Clean
			leftr <- p.left
			p.left = nil
			log.Println(p.name + " already put left")
			p.isBusy.Unlock()
			
		}
	case rightr := <- p.rightContending: // put right
		if p.right != nil && !p.right.IsClean(){
			p.isBusy.Lock()
			log.Println(p.name + " get message, ready to pass right res")
			p.right.Status = Clean
			rightr <- p.right
			p.right = nil
			log.Println(p.name + " already put right")
			p.isBusy.Unlock()
		}
	}
}

func (p *Philosopher) watingRandonTime(){
	ri := rand.Int() % 200
	time.Sleep(time.Duration(ri) * time.Millisecond)
}