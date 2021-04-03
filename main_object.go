package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func walk(pessoa Pessoa) {
	fmt.Println(pessoa.Nome, "andou")
}

func main() {

	savio := Pessoa{
		Nome:  "Savio",
		Idade: 34,
	}
	fmt.Println(savio.Nome)
	walk(savio)

}
