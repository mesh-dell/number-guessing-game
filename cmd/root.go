package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/mesh-dell/number-guessing-game/internal/game"
)

func PlayGame() error {
	var chances int

	min := 1
	max := 100

	randomNumber := rand.Intn(max-min+1) + min

	for {
		PrintWelcome()
		// read difficulty level
		var difficultyLevel string
		fmt.Printf("\nEnter your choice: ")
		fmt.Scanln(&difficultyLevel)

		difficultyLevelInt, err := strconv.Atoi(difficultyLevel)

		if err != nil || difficultyLevelInt <= 0 || difficultyLevelInt > 3 {
			return fmt.Errorf("provide a valid difficulty")
		}

		switch difficultyLevelInt {
		case 1:
			difficultyLevel = "Easy"
			chances = 10
		case 2:
			difficultyLevel = "Medium"
			chances = 5
		default:
			difficultyLevel = "Hard"
			chances = 3
		}

		var correctGuess bool = false
		fmt.Printf("\nGreat! You have selected the %s difficulty level.\n", difficultyLevel)
		fmt.Println("Let's start the game!")
		fmt.Println()

		startTime := time.Now()

		for i := 0; i < chances; i++ {
			fmt.Printf("Enter your guess: ")
			var guess string
			fmt.Scanln(&guess)
			guessInt, err := strconv.Atoi(guess)

			if err != nil {
				return fmt.Errorf("provide a valid guess")
			}

			if guessInt == randomNumber {
				elapsedTime := time.Since(startTime)
				fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", i+1)
				fmt.Println("It took you", elapsedTime)
				// check if new highScore

				highScore, err := game.CheckHighScore(difficultyLevelInt)

				if err != nil {
					return err
				}

				if i < highScore || highScore == 0 {
					fmt.Println()
					fmt.Println("Way to go! New high score:", i+1)
					game.SetNewHighScore(difficultyLevelInt, i+1)
				}
				correctGuess = true
				break
			}

			if randomNumber > guessInt {
				fmt.Printf("Incorrect! The number is greater than %d\n", guessInt)
			} else {
				fmt.Printf("Incorrect! The number is less than %d.\n", guessInt)
			}
		}

		fmt.Println()

		if !correctGuess {
			fmt.Println("Oops! You have failed to guess the correct number")
		}

		fmt.Printf("Do you want to play another round? (y/N)")
		var playAgain string
		fmt.Scanln(&playAgain)

		if strings.EqualFold(playAgain, "n") {
			return nil
		} else if strings.EqualFold(playAgain, "y") {
			fmt.Println()
			continue
		}

		return nil
	}
}

func PrintWelcome() {
	fmt.Println(
		`Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)`)
}
