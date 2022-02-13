package main

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	answer := make([]string, 0, n)
	for num := 1; num <= n; num++ {
		fizz := num%3 == 0
		buzz := num%5 == 0
		switch {
		case fizz && buzz:
			answer = append(answer, "FizzBuzz")
		case fizz:
			answer = append(answer, "Fizz")
		case buzz:
			answer = append(answer, "Buzz")
		default:
			answer = append(answer, strconv.Itoa(num))
		}
	}
	return answer
}
