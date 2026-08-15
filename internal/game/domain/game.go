package domain

import (
	"time"

	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
)

type Game struct {
	ID          uint
	QuestionIDs []uint
	Category    quizdomain.Category
	PlayerIDs   []uint
	StartTime   time.Time
}

type Player struct {
	ID      uint
	UserID  uint
	GameID  uint
	Score   uint
	Answers []PlayerAnswer
}

type PlayerAnswer struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
	Choice     quizdomain.AnswerChoice
}
