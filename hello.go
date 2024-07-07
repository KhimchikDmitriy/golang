package main

import "fmt"

func main1()  {
    //вариации записи переменных
    var hello string = "Hello"
    fmt.Println(hello)

    fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

    var (
        name string = "Den"
        age int = 27
    )
    fmt.Println(name, age)

    fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

    tom := "Tom"
    fmt.Println(tom)

    fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

    const (
        a = 1
        b
        c
        d = 3
        f
    )
    fmt.Println(a,b,c,d,f)
 }
 