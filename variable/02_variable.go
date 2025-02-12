package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var username string = "Ns"
	// fmt.Println("Hello, World!", username)
	// fmt.Printf("data type: %T\n ",username)


	// taking input from user version 1
	// var username string
	// fmt.Println("type your name and age: ")
	// fmt.Scan(&username)
	// fmt.Print("Hello, ", username ,"\n")

	// taking input from user version 2
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	username, _ := reader.ReadString('\n')

	fmt.Println("Hello, ", username)

}

