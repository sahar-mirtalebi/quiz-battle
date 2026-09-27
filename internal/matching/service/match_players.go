package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/thoas/go-funk"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/timestamp"

	presencedto "github.com/sahar-mirtalebi/quiz-battle/internal/presence/dto"

	"github.com/sahar-mirtalebi/quiz-battle/internal/category"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/dto"
)

func (s Service) MatchPlayers(ctx context.Context, req dto.MatchPlayersInput) (dto.MatchPlayersOutput, error) {
	var wg sync.WaitGroup
	for _, category := range category.AllCategories() {
		wg.Add(1)
		go s.match(ctx, category, &wg)

	}

	return dto.MatchPlayersOutput{}, nil
}

func (s Service) match(ctx context.Context, category category.Category, wg *sync.WaitGroup) {
	defer wg.Done()
	list, err := s.repo.GetWaitingListByCategory(ctx, category)
	if err != nil {
		// TODO - log error
		// TODO - update metric
		return
	}

	userIDs := make([]uint, 0, len(list))
	for _, l := range list {
		userIDs = append(userIDs, l.UserID)
	}

	presenceList, err := s.presenceClient.IsOnline(ctx, presencedto.IsOnlineInput{UserIDs: userIDs})
	if err != nil {
		// TODO - log error
		// TODO - update metric
		return
	}

	cutoff := uint64(timestamp.Add(-20 * time.Second))
	presenceUserIDs := make([]uint, 0, len(presenceList.Presences))
	for _, presence := range presenceList.Presences {
		if presence.Timestamp >= cutoff {
			presenceUserIDs = append(presenceUserIDs, presence.UserID)
		}
	}

	finalList := make([]domain.WaitingMember, 0, len(list))
	for _, member := range list {
		if funk.ContainsUInt(presenceUserIDs, member.UserID) {
			finalList = append(finalList, member)
		} else {
			// TODO - remove from waiting list
		}
	}

	for i := 0; i+1 < len(finalList); i += 2 {
		matchedUser := domain.WaitingPair{
			Category: category,
			Members:  []domain.WaitingMember{finalList[i], finalList[i+1]},
		}

		fmt.Println(matchedUser)
		//publish a new event for mu
		//remove mu user from repo
	}
}
