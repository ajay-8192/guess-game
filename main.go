package main

import (
	"math/rand"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have chances to guess the correct number based on levels.")

	for {
		fmt.Print(`
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
4. To Exit
`)
	
		var difficulty int;
		var tries int;
		fmt.Print("Enter your choice: ")
		fmt.Scan(&difficulty)
	
		switch difficulty {
		case 1:
			fmt.Println("Easy")
			tries = 10
		case 2:
			fmt.Println("Medium")
			tries = 5
		case 3:
			fmt.Println("Hard")
			tries = 3
		}

		targetNumber := rand.Intn(100) + 1
		fmt.Println("Let's start the game!")

		var guess int;
		for i := 0; i < tries; i++ {
			fmt.Print("Enter your guess: ")
			fmt.Scan(&guess)

			if guess == targetNumber {
				fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", tries + 1)
				break
			} else if guess < targetNumber {
				fmt.Printf("Incorrect! The number is less than %d.\n", guess)
			} else {
				fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
			}

			fmt.Println()
		}
		if guess != targetNumber {
			fmt.Printf("Incorrect! Unable to guess in %d changes: %d.\n", tries, targetNumber)
		}

		fmt.Print("Want to play again (Y/n): ")
		var s string;

		fmt.Scan(&s)
		if s != "Y" {
			break
		}
	}
}
