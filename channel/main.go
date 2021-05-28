package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 50)

	wg.Add(2)
	go func() {
		/*
			Duyet channel bang for...range
			for i := range ch {
			 	fmt.Println(i)
			}
		*/
		/*
			Cach 2
			for {
				if i, ok := <-ch; ok {
					fmt.Println(i)
				} else {
					break
				}
			}
		*/
		select {
		case i := <-ch1:
			fmt.Println(i)
		case j := <-ch2:
			fmt.Println(j)
		default:
			fmt.Println("Default case executed")
		}
		wg.Done()
	}()

	go func() {
		ch1 <- 42
		close(ch1)
		ch2 <- 27
		close(ch2)
		wg.Done()
	}()
	wg.Wait()
}
