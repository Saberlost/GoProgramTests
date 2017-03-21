package main

import (
	"fmt"
	"math/rand"
	"time"
)

func playGuessingGame(n int) {
	incorrectguess := true
	totguss := 0
	uniqueguesses := 0
	valuesGuessed := map[int]bool{}

	for incorrectguess {
		fmt.Println("Please give me your best guess")
		var guess int

		fmt.Scanln(&guess)
		totguss++
		_, used := valuesGuessed[guess]
		if !used {
			uniqueguesses++
		}
		valuesGuessed[guess] = true
		if guess == n {
			incorrectguess = false
			break
		}

		if guess < n {
			fmt.Println("The gussed value ", guess, " is too low")
		} else {
			fmt.Println("The gussed value ", guess, " is too high")
		}
	}
	fmt.Println("Congratulations, you managed to guess the correct number (", n, ")")
	if uniqueguesses != totguss {
		fmt.Println("But you actully guessed the same number at least twice! And you use a toatal of ", uniqueguesses, " diffrent numbers to find the correct answer")
	} else {
		fmt.Println("And you use a toatal of ", uniqueguesses, " guesses was used")
	}

}

func main() {

	fmt.Println("Welcome to the Guessing Game")
	fmt.Println("The game where your guessing the random generarted number.")
	fmt.Println("The number is between 1 and 100")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(99) + 1 // 1 - 100
	playGuessingGame(n)

}
