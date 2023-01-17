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
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
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
