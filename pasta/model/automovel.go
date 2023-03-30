package model

type Automovel struct {
	Ano    int
	Modelo string
	Placa  string
}

type Moto struct {
	Automovel
	Cilindradas int
}

type Carro struct {
	Automovel
	Potencia            int
	QuantidadesDePortas int
}
