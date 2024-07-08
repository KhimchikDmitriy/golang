package main

import "fmt"

//работа с типом функций
 
func main7() {
	fmt.Println("Вызов функции")
	new3()
	var f func(int, int) int = plus
	fmt.Println(f(12, 18))
	new3()
	x:= f(8, 7)
	fmt.Println(x)
	new3()
	f2 := plus
	fmt.Println(f2(33, 18))
	f2 = multiply
	fmt.Println(f2(33, 18))
	f3 := display
	f3("foo")
	new3()
	action(5, 5, plus)
	action(5, 5, multiply)
	new3()
	allIn()
	new3()
	f2 = selectFn(1)
	fmt.Println(f2(3, 3))
	f2 = selectFn(2)
	fmt.Println(f2(3, 3))
	f2 = selectFn(3)
	fmt.Println(f2(3, 3))
}

func new3(){

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

}

func plus(x int, y int) int {return x + y}
func multiply(x int, y int) int {return x * y}
func subtract(x int, y int) int {return x - y}
func display(message string) {fmt.Println(message)}

func action(n1 int, n2 int, operation func(int, int) int){
	result := operation(n1, n2)
	fmt.Println(result)
}

func allIn(){
	slice := []int{-2, 4, 3, -1, 7, -4, 23}
	sumOfEvens := sum(slice, isEven)
	fmt.Println(sumOfEvens)
	
	sumOfPositives := sum(slice, isPositive)
	fmt.Println(sumOfPositives)
}

func isEven(n int) bool {return n%2 == 0}
func isPositive(n int) bool {return n > 0}
func sum(numbers []int, criteria func(int) bool) int {
	result:= 0
	for _, val := range numbers {
		if criteria(val) {
			result += val
		}
	}
	return result
}

func selectFn(n int) (func(int, int) int) {
	if n == 1 {
		return plus
	} else if n == 2 {
		return subtract
	} else {
		return multiply
	}
}