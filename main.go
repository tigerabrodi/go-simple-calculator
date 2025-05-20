package main

import (
	"errors"
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

type CalculatorParams struct {
	Operator     rune
	FirstNumber  int
	SecondNumber int
}

func calculator(params CalculatorParams) {
	switch params.Operator {
	case OpAdd:
		fmt.Println("This is your result", params.FirstNumber+params.SecondNumber)
	case OpSubtract:
		fmt.Println("This is your result", params.FirstNumber-params.SecondNumber)
	case OpMultiply:
		fmt.Println("This is your result", params.FirstNumber*params.SecondNumber)
	case OpDivide:
		if params.SecondNumber == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}

		fmt.Println("This is your result", params.FirstNumber/params.SecondNumber)
	default:
		fmt.Println("This is not supported - should never happen considering all the validation logic")
	}
}

var ErrFirstNumberConversion = errors.New("could not convert first number")
var ErrSecondNumberConversion = errors.New("could not convert second number")

type GetArgNumbersParams struct {
	Operator rune
	args     []string
	signs    string
}

func getArgNumbers(params GetArgNumbersParams) (int, int, error) {
	firstNumber, err := strconv.Atoi(params.args[1])

	if err != nil {
		return 0, 0, ErrFirstNumberConversion
	}

	secondNumber, err := strconv.Atoi(params.args[2])

	if err != nil {
		return 0, 0, ErrSecondNumberConversion
	}

	return firstNumber, secondNumber, nil
}

func main() {
	signs := "/*+-"

	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Not enough arguments specified")
		return
	}

	operator := args[0]

	operatorRune, err := stringToRune(operator)

	if err != nil {
		fmt.Println("Could not convert sign to rune, you did not give us the right thing")
		return
	}

	firstNumber, secondNumber, err := getArgNumbers(GetArgNumbersParams{
		args:     args,
		signs:    signs,
		Operator: operatorRune,
	})

	if err != nil {
		if errors.Is(err, ErrFirstNumberConversion) {
			fmt.Println(err)
		} else if errors.Is(err, ErrSecondNumberConversion) {
			fmt.Println((err))
		}

		return
	}

	calculator(CalculatorParams{
		Operator:     operatorRune,
		FirstNumber:  firstNumber,
		SecondNumber: secondNumber,
	})
}
