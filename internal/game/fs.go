package game

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func ReadHighScoreData() (HighScore, error) {
	filePath := getFilePath()
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		fmt.Println("Creating new high scores file.....")

		if err != nil {
			fmt.Println("error creating file")
			return HighScore{}, err
		}

		file.Write([]byte("{}"))
		defer file.Close()

		return HighScore{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file")
		return HighScore{}, err
	}

	defer file.Close()

	var highScores HighScore
	err = json.NewDecoder(file).Decode(&highScores)

	if err != nil {
		fmt.Println("Error decoding json")
		return HighScore{}, err
	}

	return highScores, nil
}

func WriteHighScoreData(highScore HighScore) error {
	filePath := getFilePath()
	file, err := os.Create(filePath)

	if err != nil {
		fmt.Println("error creating file")
		return err
	}

	err = json.NewEncoder(file).Encode(highScore)
	defer file.Close()

	if err != nil {
		fmt.Println("error encoding json")
		return err
	}

	return nil
}

func getFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	filePath := path.Join(cwd, "highScore.json")
	return filePath
}
