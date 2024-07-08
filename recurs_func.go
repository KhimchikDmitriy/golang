package main

import "fmt"

//работа с рекурсивными функциями
 
func main() {
	fmt.Println("Вызов функции")
	new5()
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
	fmt.Println(factorial(22))
	new5()
	fmt.Println(fibonachi(4))
	fmt.Println(fibonachi(7))
	fmt.Println(fibonachi(22))
}

func new5(){

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

}

func factorial(n uint) uint{

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}

func fibonachi(n uint) uint{
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonachi(n-1) + fibonachi(n - 2)
}
