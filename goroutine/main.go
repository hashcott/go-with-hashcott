package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	lists := []string{
		"https://www.zhashkevych.com",
		"http://google.com",
		"http://youtube.com",
		"https://stackoverflow.com",
		"https://www.udemy.com/",
	}
	wg.Add(len(lists))
	for _, link := range lists {
		go checkSite(link)
	}
	// time.Sleep(10 * time.Second)
	wg.Wait()
}

func checkSite(link string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " is down")
		wg.Done()
		return
	}

	fmt.Println(link + " is OK! " + resp.Status)
	wg.Done()
}
