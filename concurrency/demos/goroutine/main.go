package main

import "fmt"

func main() {
	go sayHello()
}

func sayHello() {
	fmt.Println("hello")
}
func brian(s string) error {
	fmt.Print(s)
	return nil
}
