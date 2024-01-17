package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
	printName("Aleksei")
}

func printName(name string) {
	fmt.Printf("Name is %s\n", name)
}