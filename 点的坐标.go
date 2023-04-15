package main

import "fmt"

func main() {
	var a float64
	var b float64
	fmt.Scanf("%f %f", &a, &b)
	if a > 0 && b > 0 {
		fmt.Println("Q1")
	} else if a > 0 && b < 0 {
		fmt.Println("Q4")
	} else if a < 0 && b < 0 {
		fmt.Println("Q3")
	} else if a < 0 && b > 0 {
		fmt.Println("Q2")
	} else if a == 0 && b != 0 {
		fmt.Println("Eixo Y")
	} else if a != 0 && b == 0 {
		fmt.Println("Eixo X")
	} else if a == 0 && b == 0 {
		fmt.Println("Origem")
	}
}
