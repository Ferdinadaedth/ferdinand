package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
	fmt.Println("Please input your guess")
	for {
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		fmt.Scanln()
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
