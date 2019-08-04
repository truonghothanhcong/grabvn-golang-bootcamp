package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
	"strings"
	"strconv"
)

func eval(a float64, b float64, arithmetic string) (float64, error) {
	var result float64 = 0
	var err error = nil

	switch arithmetic {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if 0 == int(b) {
			err = errors.New("Second param must different 0")
			break
		}
		result = a / b
	default:
		err = errors.New("Invalid arithmetic")
	}

	return result, err
}

func runHomework1() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")

	scanner.Scan()
	text := scanner.Text()
	s := strings.Fields(text)
	if len(s) != 3 {
		fmt.Printf("ERROR: invalid input\n")
		return
	}

	a, err1 := strconv.ParseFloat(s[0], 64)
	if err1 != nil {
		fmt.Printf("ERROR: first element must be float type\n")
		return
	}
	b, err2 := strconv.ParseFloat(s[2], 64)
	if err2 != nil {
		fmt.Printf("ERROR: second element must be float type\n")
		return
	}

	result, err3 := eval(a, b, s[1])
	if err3 != nil {
		fmt.Printf("ERROR: %s\n", err3)
		return
	}

	msg := fmt.Sprintf("%.2f %s %.2f = %.2f\n", a, s[1], b, result)
	fmt.Printf(msg)
}