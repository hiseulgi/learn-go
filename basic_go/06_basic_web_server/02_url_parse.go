// * URL PARSING
// mengambil informasi / data pada url mulai dari protocol, host, path, query params
package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "http://kalipare.com:80/hello?name=john wick&age=27"
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	queryData := u.Query() // return map[string][]string
	name := u.Query()["name"][0]
	age := u.Query()["age"][0]
	fmt.Println(queryData)
	fmt.Printf("name: %s, age: %s\n", name, age)
}
