package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{ //slice of string having links.
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string) //creating a channel of type string,using make keyword...
	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}	

func checkLink(link string, c chan string) { //func taking link string and channel of string type.
	_, err := http.Get(link) //only checking if there is error.
	if err != nil {
		fmt.Println(link, "not working..") //printing error
		c <- link                          //passing error message to channel.
		return                             //if error occur return.
	}
	fmt.Println(link, "yes its working ") //if error doea not occurs
	c <- link                             //passing string to channel.
}
