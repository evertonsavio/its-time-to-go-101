package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

//need * to point to same memory location
func (p *Pessoa) setName(name string) {
	p.Nome = name
	fmt.Println(p.Nome)
}

func main() {

	person := Pessoa{
		Nome:  "Savio",
		Idade: 10,
	}
	person.setName("everton")
	fmt.Println(person.Nome)

}
