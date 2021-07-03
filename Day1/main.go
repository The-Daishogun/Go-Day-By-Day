package main

import "fmt"

const greetingWord = "Hello, "

func Hello(name string) string {
	return greetingWord + name
}

func main() {
	fmt.Println(Hello("The_Daishogun"))
}
