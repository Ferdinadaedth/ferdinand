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
		s := make([]int, length)
		for i := 0; i < length; i++ {
			s[i] = m + i
		}
		for _, v := range s {
			fmt.Printf("%d ", v)
			sum += v
		}
		fmt.Printf("Sum=%d\n", sum)
	}
}

