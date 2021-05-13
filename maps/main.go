package main

import "fmt"

func main() {
	// color := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }
	// fmt.Println(color)
	tester := make(map[int][]string)
	tester[1] = []string{"Nguyen", "Hanh"}
	delete(tester, 1)
	fmt.Println(tester)
}
