package game

type HighScore struct {
	Easy   int `json:"easy"`
	Medium int `json:"medium"`
	Hard   int `json:"hard"`
}

func CheckHighScore(difficulty int) (int, error) {
	highScores, err := ReadHighScoreData()

	if err != nil {
		return 0, err
	}

	switch difficulty {
	case 1:
		return highScores.Easy, nil
	case 2:
		return highScores.Medium, nil
	case 3:
		return highScores.Hard, nil
	default:
		return 0, nil
	}

}

func SetNewHighScore(difficulty, highScore int) error {
	highScores, err := ReadHighScoreData()

	if err != nil {
		return err
	}

	switch difficulty {
	case 1:
		highScores.Easy = highScore
	case 2:
		highScores.Medium = highScore
	case 3:
		highScores.Hard = highScore
	default:
		return nil
	}

	return WriteHighScoreData(highScores)
}
