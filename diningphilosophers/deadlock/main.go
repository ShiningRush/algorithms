package main

import (
	"fmt"
	"sync"

	"github.com/shiningrush/algorithms/diningphilosophers/deadlock/philosopher"
)

var mutexs []*sync.Mutex
var philosophers []*philosopher.Philosopher

func main(){
	initMutexs()
	initPhilosophers()
	for _, v := range philosophers{
		go v.Start()
	}

	g := ""
	fmt.Scanln(&g)
}

func initMutexs(){
	mutexs = append(mutexs ,&sync.Mutex{})
	mutexs = append(mutexs ,&sync.Mutex{})
	mutexs = append(mutexs ,&sync.Mutex{})
	mutexs = append(mutexs ,&sync.Mutex{})
	mutexs = append(mutexs ,&sync.Mutex{})
}

func initPhilosophers(){
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph1", mutexs[0], mutexs[4] ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph2", mutexs[1], mutexs[0] ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph3", mutexs[2], mutexs[1] ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph4", mutexs[3], mutexs[2] ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph5", mutexs[4], mutexs[3] ))
}