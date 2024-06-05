package models

type Conta struct {
	Id    int       `json: "id"`
	Nome  string    `json: "nome"`
	Valor float64   `json: "valor"`
	Tipo  TipoConta `json: "tipo"`
}

type TipoConta int

const (
	create TipoConta = iota
	Corrente
	Poupanca
)

var Contas []Conta
