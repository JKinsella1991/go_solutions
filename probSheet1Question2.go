//James Kinsella - G00282261
//Go Problem Sheet - October 2017

package main

import (
	"fmt"
	"time"
)

func main() {
	time_var := time.Now();


	fmt.Printf("The time is %d hours, %d minutes and %d seconds", time_var.Hour(), time_var.Minute(), time_var.Second())
}
