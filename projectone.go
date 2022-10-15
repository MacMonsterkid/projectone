package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	fmt.Println("Hello, world!")

	var arr []string

	arr = make([]string, 3)
	arr = append(arr, "Hello")
	arr = append(arr, "Test")
	arr = append(arr, "Good")

	for _, value := range arr {
		fmt.Println(value)
	}

	userlist := map[string]int{"Werner": 10, "Heiner": 20}

	for index, user := range userlist {
		fmt.Println(index)
		fmt.Println(user)

	}

	var p Person
	p.name = "James"
	p.age = 25
	p.job = "IT Manager"
	p.salary = 70000

	fmt.Println(p)

	//var m map[int]*Person

}
