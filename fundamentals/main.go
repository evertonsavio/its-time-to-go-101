package main

import (
	"fmt"
)

type secretAgent struct { // STRUCT

	person
	licenseToKill bool
}

type person struct { //STRUCT

	fname string
	lname string
}

func (p person) speak() {
	fmt.Println("OK THEN")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, "OPS")
}

//INTERFACE BEGIN////////////
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
//INTERFACE END//////////////

func main() {
	x := 7
	fmt.Println(x)

	xi := []int{2, 4, 5, 6} //SLICE
	fmt.Println(xi)         //HOLD LIST OF INTS

	m := map[string]int{ //MAP
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	p1 := person{
		"Miss",
		"monetpenny",
	}
	fmt.Println(p1)

	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()

	saySomething(p1)
	saySomething(sa1)

}
