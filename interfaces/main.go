package main

import "fmt"

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

func (e *Engineer) GetName() string {
	return "Name:" + e.Name
}

func PrintEmployName(e Employee) {
	fmt.Println(e.GetName())
}

func main() {
	engineer := &Engineer{
		Name: "Leandro Dev Golang",
	}
	PrintEmployName(engineer)
}
