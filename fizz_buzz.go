package main //1

import (
    "strconv"
    "fmt"
)

func main() { //3

	fmt.Println("MAIN")


}

func FizzBuzz(i int) string{

	output := "FIZZ"

	switch i {
	case 3:
		output = "FIZZ"
	case 5:
		output = "BUZZ"
	case 15:
		output = "FIZZBUZZ"
	default:
		output = strconv.Itoa(i) 
	}

	return output
}