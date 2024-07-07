package main

import "fmt"
 
func main4() {
	// синтаксис циклов
	
    for i := 1; i < 3; i++ {
		fmt.Println(i * i)
	}

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

	var b = 1
	for ; b < 3; b++{
    fmt.Println(b * b)

	}

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

	var a = 1
	for a < 3{
    fmt.Println(a * a)
    a++
	}

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")
	
	for c := 1; c < 10; c++{
		for j := 1; j < 10; j++{
			fmt.Print(c * j, "\t")
		}
		fmt.Println()
	}

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")
	
	users := [3]string{"user1", "user2", "user3"}
	for index, value := range users {
		fmt.Println(index, value)
	}

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")
	
	for _, value := range users{
		fmt.Println(value)
	}

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")
	
	var names = [3]string{"Tom", "Alice", "Kate"}
	for i:= 0; i < len(names); i++{
		fmt.Println(names[i])
	}

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")
	
	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum = 0
	
	for _, value := range numbers{
		if value < 0{
			continue        //переходим к следующему значению
		}
		sum += value
	}
	fmt.Println("Sum:", sum)

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

	var new_numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var summ = 0
	
	for _, value := range new_numbers{
		if value > 4{
			break       // если число больше 4 выходим из цикла
		}
		summ += value
	}
	fmt.Println("Sum:", summ)
}