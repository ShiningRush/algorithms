package main

import (
	"fmt"
	"github.com/shiningrush/algorithms/diningphilosophers/resourcehierarchy/philosopher"
)

var mutexs []*philosopher.HierarchicalMutex
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
	mutexs = append(mutexs ,philosopher.NewHierarchicalMutex(0))
	mutexs = append(mutexs ,philosopher.NewHierarchicalMutex(1))
	mutexs = append(mutexs ,philosopher.NewHierarchicalMutex(2))
	mutexs = append(mutexs ,philosopher.NewHierarchicalMutex(3))
	mutexs = append(mutexs ,philosopher.NewHierarchicalMutex(4))
}

func initPhilosophers(){
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph1", mutexs[0], mutexs[4] ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph2", mutexs[1], mutexs[0] ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph3", mutexs[2], mutexs[1] ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph4", mutexs[3], mutexs[2] ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph5", mutexs[4], mutexs[3] ))	
}