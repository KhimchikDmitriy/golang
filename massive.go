package main

import "fmt"

func mass() {
	var numbers [5]int = [5]int{1,2,3,4,5}
    fmt.Println(numbers[0])     
    fmt.Println(numbers[4])     
    numbers[0] = 87
    fmt.Println(numbers[0])  

    colors := [3]string{2: "blue", 0: "red", 1: "green"}
    fmt.Println(colors[2])
}