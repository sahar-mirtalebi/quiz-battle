package dto

type ProfileInput struct {
	UserID uint `json:"id"`
}

type ProfileOutput struct {
	Name string `json:"name"`
}
