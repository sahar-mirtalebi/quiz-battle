// Package backofficeuserservice contains administrative user use cases.
package backofficeuserservice

import userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"

type Repository interface {
	ListAll() ([]userdomain.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) ListAllUsers() ([]userdomain.User, error) {
	return s.repo.ListAll()
}
