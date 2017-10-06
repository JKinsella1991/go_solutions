//James Kinsella - G00282261
//Go Problem Sheet - October 2017

//Code pulled from - https://gist.github.com/abesto/3476594

package main

import (
	"fmt"
	"math"
)

const DELTA = 0.0000001
const INITIAL_Z = 100.0

func Sqrt(x float64) (z float64) {
    z = INITIAL_Z
    
    step := func() float64 {
    	return z - (z*z - x) / (2 * z)
    }
    
    for zz := step(); math.Abs(zz - z) > DELTA
    {
    	z = zz
	zz = step()
    }
    return
}

func main() {
	fmt.Println(Sqrt(500))
	fmt.Println(math.Sqrt(500))
}