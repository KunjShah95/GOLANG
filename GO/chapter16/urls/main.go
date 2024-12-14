package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning url")
	myURL := "https://example.com:8080/path/to/file?param1=value1&param2=value2"
	fmt.Printf("Type of url: %T\n", myURL)

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Can't parse URL:", err)
		return
	}
	fmt.Printf("Type of URL: %T\n", parsedURL)

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)

	//modifying url components
	parsedURL.Path = "/new/path"
	parsedURL.RawQuery = "username=iamkunj"

	//constructing  a new url string from a url object
	newURL := parsedURL.String()
	fmt.Println("New URL:", newURL)
}
