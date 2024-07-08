package main

import "fmt"

//работа с defer и panic
 
func main() {
	defer finish()
	defer fmt.Println("Finished")
	fmt.Println("Вызов функции")
	new6()
	fmt.Println(divide(15, 5))
	fmt.Println(divide(15, 0))
}

func new6(){

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

}

func finish(){
	fmt.Println("fin")
}

func divide(x,y float64) float64 {
	if y == 0 {
		panic("Ommnissia not statisfied!")
	}
	return x / y
}