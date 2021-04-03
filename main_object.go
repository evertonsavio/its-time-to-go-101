package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {

	savio := Pessoa{
		Nome:  "Savio",
		Idade: 34,
	}
	fmt.Println(savio.Nome)

}
