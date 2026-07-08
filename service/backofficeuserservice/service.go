package backofficeuserservice

import "github.com/sahar-mirtalebi/quiz-battle/entity"

type Service struct {
}

func New() Service {
	return Service{}
}

func (s Service) ListAllUser() ([]entity.User, error) {
	return []entity.User{
		{
			ID:          1,
			Name:        "fake",
			PhoneNumber: "fake",
			Role:        entity.AdminRole,
		},
	}, nil
}
