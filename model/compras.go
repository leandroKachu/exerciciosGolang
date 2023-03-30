package model

import "time"

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItensDeCompra
}

type ItensDeCompra struct {
	Nome string
}

func NewCompra(mercado string, data time.Time, nomeDosItens []string) *Compra {

	var itens []ItensDeCompra

	for _, nome := range nomeDosItens {
		itens = append(itens, ItensDeCompra{Nome: nome})
	}

	return &Compra{
		Mercado: mercado,
		Data:    data,
		Itens:   itens,
	}

}
