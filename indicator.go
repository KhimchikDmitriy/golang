package main

import "fmt"

//работа с индикаторами
 
func main14() {
	fmt.Println("Вызов функции")
	new9()
	x := 4
	p := &x
	fmt.Println("x =", x)
	fmt.Println("&x =", &x)
	fmt.Println("*p =", *p)
	fmt.Println("p =", p)
	new9()
	*p = 25
	fmt.Println("x =", x)
	new9()
	// q := new(int)
	// fmt.Println("*q =", *q)
	// *q = 12
	// fmt.Println("*q =", *q)
}

func new9(){

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

}