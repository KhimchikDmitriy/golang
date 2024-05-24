package main

import "fmt"
 
func ifElse() {
     
	c := 8
    b := 8
    if c < b {
        fmt.Println("c меньше b")
    }else if c > b{
        fmt.Println("c больше b")
    } else{
        fmt.Println("c равно b")
    }

	a := 8
    switch(a) {
        case 9: 
            fmt.Println("a = 9")
        case 8: 
            fmt.Println("a = 8")
        case 7: 
            fmt.Println("a = 7")
		default: 
            fmt.Println("значение переменной a не определено")
    }
}