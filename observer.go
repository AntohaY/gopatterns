package main

import (
	"fmt"
	"sync"
)

type (
	Observable interface {
		Add(observer Observer)
		Notify(event interface{})
		Remove(event interface{})
	}

	Observer interface {
		NotifyCallback(event interface{})
	}

	CPU struct {
		observer sync.Map
	}

	Panel struct {
		speed int
	}
)

func (wt *CPU) Add(observer Observer) {
	wt.observer.Store(observer, struct{}{})
}

func (wt *CPU) Remove(observer Observer) {
	wt.observer.Delete(observer)
}

func (wt *CPU) Notify(event interface{}) {
	wt.observer.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}

		key.(Observer).NotifyCallback(event)
		return true
	})
}

func (s Panel) NotifyCallback(event interface{}) {
	fmt.Printf("Current speed is %d km/h,  %s\n", s.speed, event)
}

func main() {
	CPU := CPU{}
	p_1 := Panel{speed: 1}
	p_2 := Panel{speed: 2}

	CPU.Add(p_1)
	CPU.Add(p_2)

	CPU.Notify("speed has been changed")

	CPU.Notify("speed has been changed")
}
