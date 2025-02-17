package main

import "fmt"

func main() {

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	for i:=0; i<len(days); i++ {
		fmt.Println(days[i])	
	}

	// another way to loop through slice
	// it gives index value in d
	for d := range days {
		fmt.Println(days[d])
	}

	// another way to loop through slice
	// it gives index and value in i and day 
	for _, day := range days {
		fmt.Println(day) // Directly prints the day name
	}
	
}