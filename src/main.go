package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	testcase := "172.16.254.1"
	fmt.Println(testcase)
	fmt.Println("is it valid IPv4?")
	fmt.Println(validIPv4Address(testcase))
}

func init() {
	fmt.Println("Welcome to init() function")
}
