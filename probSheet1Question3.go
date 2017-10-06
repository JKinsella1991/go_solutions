//James Kinsella - G00282261
//Go Problem Sheet - October 2017

package main

import "fmt"

func main() {
	// set a for loop
	for i := 1; i <= 100; i++ {

		if i % 3 == 0 {
			fmt.Printf("Buzz \n")
		}else if i % 5 == 0 {
			fmt.Printf("Fizz \n")
		}else{
			fmt.Printf("%d \n", i)
		}
	}

}
