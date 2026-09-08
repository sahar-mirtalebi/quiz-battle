package service

import (
	"context"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(ctx context.Context, req dto.RegisterInput) (dto.RegisterOutput, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterOutput{}, richerror.New().WithError(err).WithKind(richerror.KindUnexpected).WithMessage("error hashing the password")
	}

	user := userdomain.User{
		Name:           req.Name,
		PhoneNumber:    req.PhoneNumber,
		HashedPassword: string(hashedPassword),
		Role:           userdomain.UserRole,
	}

	createdUser, err := s.Repo.Register(ctx, user)
	if err != nil {
		return dto.RegisterOutput{}, richerror.New().WithError(err)
	}

	resp := dto.RegisterOutput{
		User: dto.UserInfo{
			ID:          createdUser.ID,
			Name:        createdUser.Name,
			PhoneNumber: createdUser.PhoneNumber,
		},
	}

	return resp, nil
}
