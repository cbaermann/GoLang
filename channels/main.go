package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		//go main routine makes child routines
		//main routine controlls when it exits/quits
		go checkLink(link, c)
	}
	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down"
		return
	}

	fmt.Println(link, "is working!")
	c <- "Yup it's up"

}
