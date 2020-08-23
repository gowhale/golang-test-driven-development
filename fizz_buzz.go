package main //1

import (
    "strconv"
    "fmt"
)

func main() { //3

	fmt.Println("MAIN")


}

func FizzBuzz(i int) string{

	output := ""

	if i % 15 == 0 {
		output = "FIZZBUZZ"
	} else if i % 5 == 0 {
		output = "BUZZ"
	} else if i % 3 == 0 {
		output = "FIZZ"
	} else { 
		output = strconv.Itoa(i) 
	}

	return output
}