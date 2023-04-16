package main

import "fmt"

func multiple(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * multiple(n-1)
	}
}
func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(multiple(a))
}

