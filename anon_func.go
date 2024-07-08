package main

import "fmt"

//работа с анонимными функциями
 
func main8() {
	fmt.Println("Вызов функции")
	new4()
	fu := func(x, y int) int { return x + y }
	fmt.Println(fu(8, 7))
	new4()
	newAction(10, 25, func (x, y int) int { return x + y})
	newAction(7, 5, func (x,y int) int { return x * y})
	new4()
	fu = selectFun(1)
	fmt.Println(fu(3, 3))
	fu = selectFun(2)
	fmt.Println(fu(3, 3))
	fu = selectFun(3)
	fmt.Println(fu(3, 3))
	new4()
	fu2 := square()
	fmt.Println(fu2())
	fmt.Println(fu2())
	fmt.Println(fu2())
	fmt.Println(fu2())
}

func new4(){

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

}

func newAction(n1 int, n2 int, operation func(int, int) int){
	result := operation(n1, n2)
	fmt.Println(result)
}

func selectFun(n int) (func(int, int) int){
	if n == 1 {
		return func(x,y int) int {return x + y }
	}else if n == 2 {
		return func(x,y int) int {return x * y }
	}else{
		return func(x,y int) int {return x - y }
	}
}

func square() func() int {
	var x int = 1
	return func() int {
		x++
		return x * x
	}
}