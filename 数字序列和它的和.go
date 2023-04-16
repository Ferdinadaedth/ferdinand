package main

import "fmt"

func main() {
	for {
		var m, n, temp int
		_, err := fmt.Scan(&m, &n)
		if err != nil {
			break
		}
		if m <= 0 || n <= 0 {
			break
		}
		if m > n {
			temp = m
			m = n
			n = temp
		}
		length := (n - m) + 1
		sum := 0
		for i := 0; i < length; i++ {
			sum += m + i
		}
		fmt.Printf("%d ", m)
		for i := 1; i < length; i++ {
			fmt.Printf("%d ", m+i)
		}
		fmt.Printf("sum=%d\n", sum)
	}
}
