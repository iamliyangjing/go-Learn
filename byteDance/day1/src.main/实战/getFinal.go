package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input")
			return
			continue
		}
		// \r\n
		input = strings.TrimSuffix(input, "\r\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invaild input. Please enter an integer value")
			return
			continue
		}
		fmt.Println("your guess is", guess)
		if guess > secretNumber {
			fmt.Println("your guess is bigger than the secret number,please try again")
		} else if guess < secretNumber {
			fmt.Println("smaller than")
		} else {
			fmt.Println("Corret,you legend")
			break
		}
	}

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An errror occured while readint input.pleadse try again", err)
		return
	}
	input = strings.TrimSuffix(input, "\n")

	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invail input.please enter an integer value")
		return
	}
	fmt.Println(guess)
}
