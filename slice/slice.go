package main

import (
	"fmt"
	"sort"
)

func main() {

	// here it allocate memory for 3 integer
	var sly = make([]int, 3)

	sly[0] = 1
	sly[1] = 2
	sly[2] = 3

	fmt.Println("sly: ", sly)
	
	// here it append 4 to the slice and dynamically increase the size of the slice.
	sly = append(sly, 55,33,5,99,6,9)

	fmt.Println("sly updated: ", sly)
	sort.Ints(sly)

	fmt.Println("sly sorted: ", sly)


// using splice append method to remove element from slice

    // Define a slice
    numbers := []int{10, 20, 30, 40, 50}

    // Index of the element to remove (middle element)
    index := len(numbers) / 2 // Removes the element at index 2 (30)

    // Remove the element using append
    numbers = append(numbers[:index], numbers[index+1:]...)

    // Print the updated slice
    fmt.Println(numbers) // Output: [10 20 40 50]
}

