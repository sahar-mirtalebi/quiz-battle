package domain

import (
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/internal/category"
	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
)

type Game struct {
	ID          uint
	QuestionIDs []uint
	Category    category.Category
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
