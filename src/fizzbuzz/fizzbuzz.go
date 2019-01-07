package fizzbuzz

import "strconv"

func IsFizzBuzz(number int) bool {
	return IsFizz(number) && IsBuzz(number)
}

func FizzBuzz(text string) string {
	number, _ := strconv.Atoi(text)
	if IsFizzBuzz(number) {
		return "FizzBuzz"
	}
	if IsBuzz(number) {
		return "Buzz"
	}
	if IsFizz(number) {
		return "Fizz"
	}
	return text
}
