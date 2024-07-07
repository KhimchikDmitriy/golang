package main

import "fmt"

func main2() {
    //синтаксис массива

	var numbers [5]int = [5]int{1,2,3,4,5}
    fmt.Println(numbers[0])     
    fmt.Println(numbers[4]) 
    numbers[0] = 87
    fmt.Println(numbers[0])  

    fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

    colors := [3]string{2: "blue", 0: "red", 1: "green"}
    fmt.Println(colors[2])
}