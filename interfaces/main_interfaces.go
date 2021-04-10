package main

import (
	"encoding/json"
	"fmt"
)

//STRUCT
type Car struct {
	CarName string `json:"car"`
	CarYear int    `json:"year"`
}

//INTERFACE
type vehicle interface { start() string}

//INTERFACE IMPLEMENTATION
func (c Car) start() string{
	return "Started"
}

//FUNCTION THAT RECEIVES A CAR THAT IMPLEMENTS VEHICLE
func example(car vehicle){
	fmt.Println("OOPS, its a vehicle!")
}

//STRUCT METHOD WITH AN DUMMY SETYEAR
func (c *Car) drive() {
	c.CarYear = 2021
	fmt.Println(c.CarName, "Andou!")
	fmt.Printf("%T \n", c.CarYear)
	fmt.Printf("%v \n", c.CarYear)
}

//MAIN FUNC
func main() {

	car1 := Car{
		CarName: "Fusion",
		CarYear: 2020,
	}

	fmt.Println(car1.CarName)
	car1.drive()

	//TO JSON FROM STRUCT
	result, _ := json.Marshal(car1)
	fmt.Println(string(result))

	//TO STRUCT FROM JSON
	j := []byte(`{"car":"BMW", "year":2022}`)
	var car2 Car
	json.Unmarshal(j, &car2)
	fmt.Println(car2.CarName)

	//Interfaces example
	example(car2)

}
