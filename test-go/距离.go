package main

import "fmt"

func main() {
	var l float64
	fmt.Scan(&l)
	time := (l / 30) * 60
	fmt.Printf("%.0f minutos\n", time)
}
