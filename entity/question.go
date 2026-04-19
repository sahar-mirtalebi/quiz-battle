package entity

type Question struct {
	ID             uint
	content        string
	PossibleAnswer []PossibleAnswer
	// this is possibleAnswerID
	CorrctAnswerID uint
	Difficulty     QuestionDifficulty
	CategoryID     uint
}

type PossibleAnswer struct {
	ID      uint
	content string
	choice  PossibleAnswerChoise
}

type PossibleAnswerChoise uint8

func (p PossibleAnswerChoise) IsValid() bool {
	if p >= PossibleAnswerA && p <= PossibleAnswerD {
		return true
	}
	return false
}

const (
	PossibleAnswerA PossibleAnswerChoise = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)

type QuestionDifficulty uint8

func (p QuestionDifficulty) IsValid() bool {
	if p >= QuestionDifficultyEasy && p <= QuestionDifficultyHard {
		return true
	}
	return false
}

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)
