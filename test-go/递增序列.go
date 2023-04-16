package main

import "fmt"

func main() {
	for {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		for i := 1; i <= x; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
	}
}
