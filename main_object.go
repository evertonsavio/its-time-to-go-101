package main

import (
	"errors"
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func walk0(pessoa Pessoa) {
	fmt.Println(pessoa, "andouuu")
}
func walk(pessoa Pessoa) error {
	if pessoa.Nome != "Savio" {
		return errors.New("Nome deve ser Savio")
	}
	fmt.Println(pessoa.Nome, "andou")
	return nil
}
func walk2(pessoa Pessoa) (string, error) {
	if pessoa.Nome != "Savio" {
		return "", errors.New("Nome deve ser Savio")
	}
	return pessoa.Nome + "andou", nil
}

func main() {

	savio := Pessoa{
		Nome:  "Savio",
		Idade: 34,
	}
	fmt.Println(savio.Nome)
	res, err := walk2(savio)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)

}
