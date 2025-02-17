package main

import ( "fmt"
"net/url"

)

const myUrl = "https://www.loc.dev:3000/learn?course=golang&module=web"

func main() {

res , err := url.Parse(myUrl)

if err != nil {	
	fmt.Println(err)
	return }

	fmt.Println("Scheme: ", res.Scheme )
	fmt.Println("Scheme: ", res.Host )
	fmt.Println("Scheme: ", res.Port())
	fmt.Println("Scheme: ", res.Path )
	
	queryparams := res.Query()

	fmt.Println("Query Params: ", queryparams["course"])
	fmt.Println("Query Params: ", queryparams["module"])


	for k, v :=range queryparams {
		fmt.Println(k, v)
	}

	partOfUrl := &url.URL{
		Scheme: "https",
		Host: "www.loc.dev:3000",	
		Path: "/learn",
		RawQuery: "course=golang&module=web",
	}
		fmt.Println(partOfUrl.String())
}