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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		PrintWelcome()
		level, chances, err := chooseDifficulty()
		if err != nil {
			fmt.Println(err)
			continue
		}

		difficulty := map[int]string{1: "Easy", 2: "Medium", 3: "Hard"}
		fmt.Printf("\nGreat! You have selected the %s difficulty level.\n", difficulty[level])
		fmt.Println("Let's start the game!")
		fmt.Println()

		target := r.Intn(100) + 1
		start := time.Now()
		win, attempts := playRound(target, chances)

		if win {
			elapsedTime := time.Since(start)
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			fmt.Println("It took you", elapsedTime)
			updateHighScore(level, attempts)
		} else {
			fmt.Println("Out of chances! Better luck next time.")
		}

		fmt.Println()
		fmt.Printf("Do you want to play another round? (y/N)")
		var playAgain string
		fmt.Scanln(&playAgain)

		if !strings.EqualFold(playAgain, "y") {
			return nil
		}
		fmt.Println()
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

func chooseDifficulty() (int, int, error) {
	fmt.Printf("\nChoose difficulty: ")
	var input string
	fmt.Scanln(&input)
	level, err := strconv.Atoi(input)
	if err != nil || level < 1 || level > 3 {
		return 0, 0, fmt.Errorf("invalid difficulty")
	}
	chances := map[int]int{1: 10, 2: 5, 3: 3}[level]
	return level, chances, nil
}

func playRound(target, chances int) (bool, int) {
	for i := 1; i <= chances; i++ {
		fmt.Printf("Enter your guess: ")
		var input string
		fmt.Scanln(&input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Try again")
			continue
		}

		switch {
		case guess == target:
			return true, i
		case guess < target:
			fmt.Printf("Incorrect! The number is greater than %d\n", guess)
		default:
			fmt.Printf("Incorrect! The number is less than %d.\n", guess)
		}
	}
	return false, 0
}

func updateHighScore(level, attempts int) {
	high, err := game.CheckHighScore(level)
	if err != nil {
		fmt.Println("Error checking high score")
	}
	if high == 0 || attempts < high {
		fmt.Println("🎉 New high score:", attempts)
		game.SetNewHighScore(level, attempts)
	}
}
