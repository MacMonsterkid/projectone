package main

import "fmt"

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
}
