package service

import (
	"context"

	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/dto"
	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func (s Service) MatchPlayers(ctx context.Context, req dto.MatchPlayersInput) (dto.MatchPlayersOutput, error) {
	pairs := make([]domain.WaitingPair, 0)
	for _, category := range quizdomain.AllCategories() {
		list, err := s.repo.GetWaitingListByCategory(ctx, category)
		if err != nil {
			return dto.MatchPlayersOutput{}, richerror.New().WithError(err).
				WithMeta(map[string]interface{}{"category": category})
		}
		for i := 0; i+1 < len(list); i += 2 {
			pairs = append(pairs, domain.WaitingPair{
				Category: category,
				Members:  []domain.WaitingMember{list[i], list[i+1]},
			})
		}
	}

	return dto.MatchPlayersOutput{Pairs: pairs}, nil
}
