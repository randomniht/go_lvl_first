package main

import "fmt"

func main_new() {
	var (
		name string = "Tom"
		age  int    = 27
	)

	test := "new"
	fmt.Println(test)
	fmt.Println(name)
	fmt.Println(age)
}

func main() {
	fmt.Println("commit test")
	main_new()
}
