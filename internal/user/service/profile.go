package userservice

import (
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"
)

func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {

	user, err := s.Repo.GetUserByID(req.UserID)
	if err != nil {
		return param.ProfileResponse{}, richerror.New().WithError(err).WithMeta(map[string]interface{}{"req": req})
	}
	return param.ProfileResponse{Name: user.Name}, nil
}
