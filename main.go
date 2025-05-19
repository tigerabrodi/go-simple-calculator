package main

import (
	"fmt"
	"os"
	"strconv"
)

func stringToRune(s string) (rune, error) {
	if len([]rune(s)) != 1 {
		return 0, fmt.Errorf("input must be a single character, got %d characters", len([]rune(s)))
	}

	r := []rune(s)[0]

	if r != '/' || r != '+' || r != '-' || r != '*' {
		return 0, fmt.Errorf("The sign you specified was not one of the allowed character. We got %c", r)
	}

	return r, nil
}

func main() {
	signs := "/*+-"
	signsSet := make(map[rune]struct{}, len(signs))

	for _, sign := range signs {
		signsSet[sign] = struct{}{}
	}

	args := os.Args[1:]

	signArg := args[0]

	isOperator := func(char rune) bool {
		_, exists := signsSet[char]
		return exists
	}

	signAsRune, err := stringToRune(signArg)

	if err != nil {
		fmt.Printf("what we got %c as sign was not what we expected, we expect to get one of these, only one of them %s", signAsRune, signs)
		return
	}

	if !isOperator(signAsRune) {
		fmt.Println("the letter we got is not one of the signs")
		return
	}

	firstNumber, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Printf("Could not convert %s to a number", args[1])
		return
	}

	secondNumber, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Printf("We could not convert the second number to number, what we got was %v", secondNumber)
		return
	}

	switch signAsRune {
	case '+':
		fmt.Println(firstNumber + secondNumber)
	case '-':
		fmt.Println(firstNumber - secondNumber)
	case '*':
		fmt.Println(firstNumber * secondNumber)
	case '/':
		if secondNumber == 0 {
			panic("division by zero")
		}

		fmt.Println(firstNumber / secondNumber)
	}
}
