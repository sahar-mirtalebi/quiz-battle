package matchingservice

import (
	"context"
	"testing"
	"time"

	param "github.com/sahar-mirtalebi/quiz-battle/internal/matching/dto"
	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
)

type matchingRepoStub struct {
	userID   uint
	category quizdomain.Category
}

func (r *matchingRepoStub) AddToWaitingList(_ context.Context, userID uint, category quizdomain.Category) error {
	r.userID = userID
	r.category = category
	return nil
}

func TestAddToWaitingList(t *testing.T) {
	repo := &matchingRepoStub{}
	waitingTimeout := 30 * time.Second
	service := New(repo, Config{WaitingTimeout: waitingTimeout})

	response, err := service.AddToWaitingList(context.Background(), param.AddToWaitingListRequest{
		UserID:   42,
		Category: quizdomain.FootballCategory,
	})
	if err != nil {
		t.Fatal(err)
	}
	if repo.userID != 42 || repo.category != quizdomain.FootballCategory {
		t.Fatalf("unexpected repository call: user=%d category=%s", repo.userID, repo.category)
	}
	if response.Timeout != waitingTimeout {
		t.Fatalf("unexpected timeout: %s", response.Timeout)
	}
}
