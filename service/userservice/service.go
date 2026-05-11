package userservice

import (
	"fmt"

	"github.com/sahar-mirtalebi/quiz-battle/dto"
	"github.com/sahar-mirtalebi/quiz-battle/entity"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/phonenumber"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/richerror"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}

type AuthService interface {
	CreateAccessToken(userID uint) (string, error)
	CreateRefreshToken(userID uint) (string, error)
}

type Service struct {
	Auth AuthService
	Repo Repository
}

func New(repo Repository, auth AuthService) Service {
	return Service{Repo: repo, Auth: auth}
}

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	// TODO - we should verify phone number by verification code
	if !phonenumber.IsValid(req.PhoneNumber) {
		return dto.RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	if isUnique, err := s.Repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return dto.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		}

		if !isUnique {
			return dto.RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	if len(req.Name) < 3 {
		return dto.RegisterResponse{}, fmt.Errorf("name should be grater than 3")
	}

	// TODO - check the password with regex
	if len(req.Password) < 8 {
		return dto.RegisterResponse{}, fmt.Errorf("password should be grater than 8")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("error hashing the password")
	}

	user := entity.User{
		Name:           req.Name,
		PhoneNumber:    req.PhoneNumber,
		HashedPassword: string(hashedPassword),
	}

	createdUser, err := s.Repo.Register(user)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	resp := dto.RegisterResponse{
		User: dto.UserInfo{
			ID:          createdUser.ID,
			Name:        createdUser.Name,
			PhoneNumber: createdUser.PhoneNumber,
		},
	}

	return resp, nil
}

func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	const op = "userservice.Login"
	// TODO - it would be better to use two seperate method for existance check and GetUserByPhoneNumber
	user, exist, err := s.Repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return dto.LoginResponse{}, richerror.New(op).WithError(err)
	}

	if !exist {
		return dto.LoginResponse{}, fmt.Errorf("phone number or password is not correct")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("phone number or password is not correct")
	}

	accessToken, err := s.Auth.CreateAccessToken(user.ID)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	refreshToken, err := s.Auth.CreateRefreshToken(user.ID)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return dto.LoginResponse{
		Tokens: dto.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		User: dto.UserInfo{
			ID:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}

func (s Service) Profile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	const op = "userservice.Profile"

	user, err := s.Repo.GetUserByID(req.UserID)
	if err != nil {
		return dto.ProfileResponse{}, richerror.New(op).WithError(err).Withmeta(map[string]interface{}{"req": req})
	}
	return dto.ProfileResponse{Name: user.Name}, nil
}
