package main

import (
	"fmt"
	"os"
)

func main() {
	// Open or create a file
	file, err := os.OpenFile("/Users/priyanka/temp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}
