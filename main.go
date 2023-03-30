package main

import (
	"exercicio1/model"
	"fmt"
	"time"
)

func main() {

	var nomeDosItens []string
	nomeDosItens = append(nomeDosItens, "feijao")
	nomeDosItens = append(nomeDosItens, "arroz")
	nomeDosItens = append(nomeDosItens, "batata")

	compra := model.NewCompra("mercadinho do pai", time.Now(), nomeDosItens)

	fmt.Println(compra)
}
