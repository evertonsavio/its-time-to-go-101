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
func walk1(pessoa Pessoa) error {
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

//Attach method to Struct
func (p Pessoa) walk3() (string, error) {
	if p.Nome != "Savio" {
		return "", errors.New("Nome deve ser Savio")
	}
	return p.Nome + "andou", nil
}

func main() {

	savio := Pessoa{
		Nome:  "Savio",
		Idade: 34,
	}
	fmt.Println(savio.Nome)
	res, err := walk2(savio)
	res, err = savio.walk3()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)

}
