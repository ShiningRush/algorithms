package main

import (
	"fmt"
)

type Rabbit struct {
	age int
}

func (r *Rabbit) Born() *Rabbit {
	r.age++
	if r.age > 1 {
		// if r.age%3 == 0 {
		return &Rabbit{}
		// }
	}

	return nil
}

func main() {
	total := []*Rabbit{&Rabbit{}}
	for i := 1; i < 15; i++ {
		for _, aRabbit := range total {
			if newRabbit := aRabbit.Born(); newRabbit != nil {
				total = append(total, newRabbit)
			}
		}

		fmt.Println(len(total))
	}
}
