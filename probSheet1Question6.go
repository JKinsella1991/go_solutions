//James Kinsella - G00282261
//Go Problem Sheet - October 2017

//Code pulled from - https://play.golang.org/p/vhHmjhOMEo

package main

import "fmt"

func main() {
	//For this program, the numbers are hardcoded in
	//Utilizing the rand generator from the previous program
	//could allow for a more random progression
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	//Initializing the comparison, comparing each numbers against
	//the current percieved largest
	smallest, biggest := x[0], x[0]
	for _, v := range x {
		if v > biggest {
			biggest = v
		}
		if v < smallest {
			smallest = v
		}
	}

	//Spits out result

	fmt.Println("The biggest number is ", biggest)
	fmt.Println("The smallest number is ", smallest)
}
