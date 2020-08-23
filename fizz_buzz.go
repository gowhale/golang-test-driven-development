package main //1

import "fmt" //2
//import "os"

func main() { //3

	fmt.Println("MAIN")

}

func FizzBuzz(i int) string{

	output := "FIZZ"

	if i == 5 {
		output = "BUZZ"
	}

	if i == 15 {
		output = "FIZZBUZZ"
	}

	if i == 1 {
		output = "1"
	}

	if i == 2 {
		output = "2"
	}


	return output
}