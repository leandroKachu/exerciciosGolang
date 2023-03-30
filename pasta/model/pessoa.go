package model

import "time"

type Pessoa struct {
	Nome           string
	Idade          int
	Endereco       Endereco
	DataNascimento time.Time
}

// Criar funcao
func CalcularIdade(p Pessoa) int {
	anoNascimento := p.DataNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoNascimento
}

// Criar metodo da struct Pessoa

func (p Pessoa) IdadeAtual() int {
	anoNascimento := p.DataNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoNascimento
}
