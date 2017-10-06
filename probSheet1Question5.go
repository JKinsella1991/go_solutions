//James Kinsella - G00282261
//Go Problem Sheet - October 2016

//Uses code from https://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html

package main

import "fmt"
import "math/rand"
import "time"

func randGen(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {

	var myname string
    myrand := randGen(1, 100)
    guessTaken := 0
    var guess int

	fmt.Printf("Welcome to my random number guessing game! \n")
	fmt.Printf("In a moment you can pick a number between 1 and 100. \n")

	fmt.Printf("I've generated a number! Can you guess what it is? \n")
	fmt.Printf("You have 15 Chances, good luck :) \n")

	for guessTaken < 15 {
        fmt.Println("Take a guess...")
        fmt.Scanf("%d", &guess)
        guessTaken++
        if guess < myrand {
            fmt.Println("Your guess is too low.")
        }else if guess > myrand {
            fmt.Println("Your guess is too high.")
        }else if guess == myrand {
            break
        }
    }
    if guess == myrand {
        fmt.Printf("Good job %s! You guessed my number in %d tries\n", myname, guessTaken)
    } else {
        fmt.Printf("Nope. The number I had in mind was %d\n", myrand)
    }
}