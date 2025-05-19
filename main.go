package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	OpAdd      = '+'
	OpSubtract = '-'
	OpMultiply = '*'
	OpDivide   = '/'
)

func stringToRune(s string) (rune, error) {
	if len([]rune(s)) != 1 {
		return 0, fmt.Errorf("input must be a single character, got %d characters", len([]rune(s)))
	}

	r := []rune(s)[0]

	if r != OpDivide && r != OpAdd && r != OpSubtract && r != OpMultiply {
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

	if len(args) < 3 {
		fmt.Println("Not enough arguments specified")
		return
	}

	signArg := args[0]

	signAsRune, err := stringToRune(signArg)

	if err != nil {
		fmt.Printf("what we got %c as sign was not what we expected, we expect to get one of these, only one of them %s", signAsRune, signs)
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
	case OpAdd:
		fmt.Println("This is your result", firstNumber+secondNumber)
	case OpSubtract:
		fmt.Println("This is your result", firstNumber-secondNumber)
	case OpMultiply:
		fmt.Println("This is your result", firstNumber*secondNumber)
	case OpDivide:
		if secondNumber == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}

		fmt.Println("This is your result", firstNumber/secondNumber)
	}
}
