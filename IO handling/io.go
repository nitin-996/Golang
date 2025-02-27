package main

import (
	"fmt"
	"strings"
)

func main() {
	statement1 := strings.NewReader("This is the Golang practice") 
	buffer := make([]byte, 4) 

	for {
		n, err := statement1.Read(buffer)  // Read up to 4 bytes
		fmt.Println(err)
		if err != nil {
			break // Stop when end of file is reached
		}
		fmt.Println(buffer[:n])       // Print byte slice (raw bytes)
		fmt.Println(string(buffer[:n])) // Convert to string and print
	}
	
}
