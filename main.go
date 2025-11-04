package main

import (
	"fmt"
	"os"

	"github.com/mesh-dell/number-guessing-game/cmd"
)

func main() {
	if err := cmd.PlayGame(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
