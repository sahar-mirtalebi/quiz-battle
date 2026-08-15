// Package dto contains the user transport DTOs.
package dto

type LoginInput struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginOutput struct {
	User   UserInfo `json:"user"`
	Tokens Tokens   `json:"tokens"`
}
