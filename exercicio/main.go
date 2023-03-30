package main

import (
	"fmt"
	"gohan/model"
	"time"
)

func main() {

	newItem := model.Item{
		NomeProduto: "",
	}

	fmt.Println(newItem.CriarItem("banana", time.Date(1999, 04, 16, 0, 0, 0, 0, time.Local)))

	fmt.Println("item:", newItem.NomeProduto, "data:", newItem.Dia)
}
