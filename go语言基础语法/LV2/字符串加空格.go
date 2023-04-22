package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("获取失败")
	}
	var result string
	for i := 0; i < len(s); i++ {
		result += string(s[i]) + " "
	}
	fmt.Println(result)
}
