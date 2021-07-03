package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	c := make(chan string)
	lists := []string{
		"https://www.zhashkevych.com",
		"http://google.com",
		"http://youtube.com",
		"https://stackoverflow.com",
		"https://www.udemy.com/",
	}
	for _, link := range lists {
		go checkSite(c, link)
	}
	for {
		go checkSite(c, <-c)
	}
	// time.Sleep(10 * time.Second)
}

func checkSite(c chan string, link string) {
	time.Sleep(5 * time.Second)
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " is down")
		c <- link
		return
	}

	fmt.Println(link + " " + resp.Status)
	c <- link
}
