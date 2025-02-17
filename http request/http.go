package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://www.google.com"

func main() {

	

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("type of url %T\n", response)
	
	// when we open a request it is our resposibility to close it as well.
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
