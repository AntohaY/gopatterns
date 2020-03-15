package main

import "fmt"

type Donkey struct {
	Name string
}

var dk *Donkey

func GetInstance() *Donkey {
	if dk == nil {
		dk = &Donkey{"Donkey"}
	}
	return dk
}

func (d Donkey) GetInfo() {
	fmt.Println(d.Name)
}

func main() {
	GetInstance().GetInfo()
	GetInstance().GetInfo()
	GetInstance().GetInfo()
}
