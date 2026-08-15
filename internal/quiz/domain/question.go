package domain

type Question struct {
	ID              uint
	Content         string
	PossibleAnswers []PossibleAnswer
	CorrectAnswerID uint
	Difficulty      QuestionDifficulty
	CategoryID      uint
}

type PossibleAnswer struct {
	ID      uint
	Content string
	Choice  AnswerChoice
}

type AnswerChoice uint8

func (p AnswerChoice) IsValid() bool {
	if p >= PossibleAnswerA && p <= PossibleAnswerD {
		return true
	}
	return false
}

const (
	PossibleAnswerA AnswerChoice = iota + 1
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
