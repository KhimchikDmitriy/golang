package main

import "fmt"

//работа с индикаторами и функциями
 
func main() {
	fmt.Println("Вызов функции")
	new10()
	d := 5
	fmt.Println("d= ", d)
	changeValue(d)
	fmt.Println("d= ", d)
	fmt.Println("but")
	fmt.Println("d= ", d)
	changeValueSecond(&d)
	fmt.Println("d= ", d)
	// new10()
	// p1 := createPointer(8)
	// fmt.Println(p1)
	// p2 := createPointer(12)
	// fmt.Println(p2)
	// p3 := createPointer(14)
	// fmt.Println(p3)
}

func changeValue(x int) {
	x = x * x
}
func changeValueSecond(x *int) {
	*x = (*x) * (*x)
}
// func createPointer(x int) *int {
// 	p := new(int)
// 	*p = x
// 	return p
// }

func new10(){

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

}