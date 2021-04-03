package main

import "fmt"

type nomeType []string

func main() {
	fmt.Println("Olá mundo!")

	var nome string //declarando
	nome = "Sávio"  // atribuindo

	nome_2 := "Sávio"  //atribuindo e declarando; Inferencia de TIPO.
	nome_2 = "EVERTON" //atrinbuindo

	var nome_list nomeType
	nome_list = append(nome_list, "SSS")
	nome_list = append(nome_list, "AAA")

	print(nome)
	print(nome_2)
	print(nome_list[0])

	for n := range nome_list {
		print(nome_list[n])
	}
	for _, v := range nome_list {
		print(v)
	}

	for i := 0; i < 10; i++ { // UNICO LACO DE REPETICAO
		print(i)
	}

}
