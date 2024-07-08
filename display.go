package main

import "fmt"

//работа с отображением
 
func main() {
	fmt.Println("Вызов функции")
	new8()
	people := map[string]int{
		"Tom": 1,
		"Bob": 2,
		"Sam": 4,
		"Alice": 8,
	}
	fmt.Println(people)
	fmt.Println(people["Bob"])
	people["Bob"] = 32
	fmt.Println(people["Bob"])
	people["Bob"] = 2
	new8()
	if val, ok := people["Alice"]; ok {fmt.Println(val)}
	new8()
	for key, value := range people {fmt.Println(key, value)}
	new8()
	fmt.Println(people)
	people["Kate"] = 220
	people["Katee"] = 220
	fmt.Println(people)
	new8()
	delete(people, "Katee")
	fmt.Println(people)
}

func new8(){

	fmt.Println(" ")
	fmt.Println("слеующий пример:")
	fmt.Println(" ")

}