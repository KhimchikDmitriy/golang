package main

import "fmt"

//синтаксис возврата
 
func main6() {
	fmt.Println("Вызов функции")
	new2()
	a := addfunc(4, 5)   
    b := addfunc(20, 6)  
    fmt.Println(a)
    fmt.Println(b)
	new2()
	c := addfunc2(17, 18) 
    fmt.Println(c)
	new2()
	age, name := addfunc3(4, 5, "Tom", "Jerson")	
	fmt.Println(age)    
    fmt.Println(name)   
}

func new2(){

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

}

func addfunc(x, y int) int {
    return x + y
}

func addfunc2(x, y int) (z int) {
    z = x + y
    return
}

func addfunc3(x, y int, firstName, lastName string) (z int, fullName string) {
    z = x + y
    fullName = firstName + " " + lastName
    return
}