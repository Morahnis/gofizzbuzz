package main

import "fmt"
import "strconv"

func main() {
	var result = ""
	for i := 1; i <= 100; i++ {
		var fb = false
		if i %3 == 0 {
			result += " Fizz"
			fb = true
		}
		if i %5 == 0 {
			result += " Buzz"
			fb = true
		}
		if fb != true {
			result += " "
			result += strconv.Itoa(i)
		}
		result += ","
	}
	fmt.Println(result)
}