package main

import "fmt"

//работа с функциями
 
func main5() {
	fmt.Println("Вызов функции")
	new()
	add(4, 5)
	new()
	add2(1, 2, 3.4, 5.6, 1.2)
	new()
	add3(1, 2, 3)        
    add3(1, 2, 3, 4)     
    add3(5, 6, 7, 2, 3)  
}

func new(){

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

}

func add(x int, y int){
    var z = x + y
    fmt.Println("x + y = ", z)
}

func add2(x, y int, a, b, c float32){
    var z = x + y
    var d = a + b + c
    fmt.Println("x + y = ", z)
    fmt.Println("a + b + c = ", d)
}

func add3(numbers ...int){
    var sum = 0
    for _, number := range numbers{
        sum += number
    }
    fmt.Println("sum = ", sum)
}