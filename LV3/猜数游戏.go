package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
                // 输入我们猜的数字
		guess, err := strconv.Atoi(string(data)) //string to int,并作输入格式判断
		// Go语言中处理错误的方法
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			return
		}
		fmt.Println("Your guess is", guess)
		if guess < secretNumber {
			fmt.Println("too small")
		} else if guess > secretNumber {
			fmt.Println("too high")
		} else {
			fmt.Println("You are right")
			break
		}
	}
}

