package model

import "time"

type Item struct {
	DataDeCompra
	NomeProduto string
}

type DataDeCompra struct {
	Dia time.Time
}

func (i *Item) CriarItem(param string, param2 time.Time) string {
	i.NomeProduto = param
	i.Dia = param2
	return i.NomeProduto
}
