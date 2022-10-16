package main

import (
	"fmt"

	"github.com/MacMonsterkid/projectone/pkg/helpers"
)

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

	var p1 Person
	p1.name = "Fritz"
	p1.age = 34
	p1.job = "Developer"
	p1.salary = 45000

	var personlist []Person
	personlist = append(personlist, p, p1)

	fmt.Println(personlist, len(personlist))

	myPort := helpers.Port

	fmt.Println(fmt.Sprintf("Listen on localhost port %s", myPort))

}
