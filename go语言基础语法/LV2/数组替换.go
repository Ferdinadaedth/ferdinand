package main

import "fmt"

func main() {
	var x [10]int
	for i := 0; i < len(x); i++ {
		fmt.Scan(&x[i])
	}
	for i := 0; i < 10; i++ {
		if x[i] <= 0 {
			x[i] = 1
		}
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("X[%d] = %d\n", i, x[i])
	}
}
