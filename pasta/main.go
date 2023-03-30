package main

import (
	"fmt"
	"projeto/model"
	"time"
)

func main() {

	endereco := model.Endereco{
		Rua:    "claudinho",
		Bairro: "Bairro",
		Numero: "Número",
	}
	pessoa := model.Pessoa{
		Nome:           "Leandro",
		Idade:          23,
		Endereco:       endereco,
		DataNascimento: time.Date(1999, 04, 16, 0, 0, 0, 0, time.Local),
	}

	automovelMoto := model.Automovel{
		Ano:    2019,
		Modelo: "Ducati",
		Placa:  "FTX3299",
	}

	moto := model.Moto{
		Automovel:   automovelMoto,
		Cilindradas: 1999,
	}

	fmt.Println("endereco:", endereco)
	fmt.Println("pessoa:", pessoa)

	// Acessando um metodo especifico de pessoa
	fmt.Println(pessoa.IdadeAtual())
	// Acessando uma funcao do model/pessoa

	fmt.Println(model.CalcularIdade(pessoa))

	// Criando heranca entre moto e automovel
	fmt.Println(automovelMoto)
	fmt.Println(moto)

}
