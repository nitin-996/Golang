package main

import "fmt"

func main() {
        var i, j = 10, 50

        switch {
        case i+j == 60:
                fmt.Println("equal to 60")
        case i+j <= 60:
                fmt.Println("less than or equal to 60")
                fallthrough
        default:
                fmt.Println("greater than 60")
        }

}

// Note: When using switch with conditionals - we don't specify expression after switch