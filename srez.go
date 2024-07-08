package main

import "fmt"

//работа со срезами
 
func main11() {
	fmt.Println("Вызов функции")
	new7()
	users := []string{"a", "b", "c", "d", "e"}
	fmt.Println(users)
	fmt.Println()
	fmt.Println(users[1])
	fmt.Println()
	users[1] = "Bob"
	fmt.Println(users[1])
	fmt.Println()
	new7()
	for _, user := range users {
		fmt.Println(user)
		fmt.Println()
	}
	new7()
	var tools []string = make([]string, 3)
	tools[0] = "t1"
	tools[1] = "t2"
	tools[2] = "t3"
	fmt.Println(tools)
	new7()
	fmt.Println(users)
	users = append(users, "user")
	fmt.Println(users)
	new7()
	users1 := users[2:5]
	users2 := users[:3]
	users3 := users[3:]
	fmt.Println(users1)
	fmt.Println(users2)
	fmt.Println(users3)
	new7()
	n := 2
	users = append(users[:n], users[n+1:]...)
	fmt.Println(users)
}

func new7(){

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

}