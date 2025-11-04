# number-guessing-game

A simple CLI number guessing game written in Go.

## Overview

Pick a difficulty, guess a number between 1–100, and try to beat your best score.

* **Easy:** 10 chances
* **Medium:** 5 chances
* **Hard:** 3 chances

High scores are stored locally in a JSON file.

## Structure

```
.
├── cmd/            # CLI logic
├── internal/game/  # Game + high score logic
├── main.go         # Entrypoint
└── go.mod
```

## Run

```bash
git clone https://github.com/mesh-dell/number-guessing-game
cd number-guessing-game
go run main.go
```

## How to Play

1. Select difficulty.
2. Guess until correct or out of tries.
3. Beat your previous high score.
4. Play again or exit.

https://roadmap.sh/projects/number-guessing-game
